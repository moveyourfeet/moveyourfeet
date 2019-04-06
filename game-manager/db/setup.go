package db

// SetupProcedures adds functions to PostgreSQLs
func SetupProcedures() {
	DB.Exec(`
CREATE OR REPLACE FUNCTION id_to_alpha(n int) RETURNS text
LANGUAGE plpgsql IMMUTABLE STRICT AS $$
DECLARE
 alphabet text:='abcdefghijklmnopqrstuvwxyz012345678';
 sign_char char:='9';
 base int:=length(alphabet); 
 _n bigint:=abs(n);
 output text:='';
BEGIN
 LOOP
   output := output || substr(alphabet, 1+(_n%base)::int, 1);
   _n := _n / base;
   EXIT WHEN _n=0;
 END LOOP;
 RETURN CASE WHEN (n<0) THEN output || sign_char::text ELSE output END;
 RETURN output;
END; $$;

CREATE OR REPLACE FUNCTION skip32(val int4, cr_key bytea, encrypt bool) returns int4
AS $$
DECLARE
	kstep int;
	k int;
	wl int4;
	wr int4;
	g1 int4;
	g2 int4;
	g3 int4;
	g4 int4;
	g5 int4;
	g6 int4;
	ftable bytea:='\xa3d70983f848f6f4b321157899b1aff9e72d4d8ace4cca2e5295d91e4e3844280adf02a017f1606812b77ac3e9fa3d5396846bbaf2639a197caee5f5f7166aa239b67b0fc193811beeb41aead0912fb855b9da853f41bfe05a58805f660bd89035d5c0a733066569450094566d989b7697fcb2c2b0fedb20e1ebd6e4dd474a1d42ed9e6e493ccd4327d207d4dec7671889cb301f8dc68faac874dcc95d5c31a47088612c9f0d2b8750825464267d0340344b1c73d1c4fd3bccfb7fabe63e5ba5ad04239c145122f02979717eff8c0ee20cefbc72756f37a1ecd38e628b8610e8087711be924f24c532369dcff3a6bbac5e6ca9135725b5e3bda83a0105592a46';
BEGIN
	IF (octet_length(cr_key)!=10) THEN
	RAISE EXCEPTION 'The encryption key must be exactly 10 bytes long.';
	END IF;
	
	IF (encrypt) THEN
	kstep := 1;
	k := 0;
	ELSE
	kstep := -1;
	k := 23;
	END IF;
	
	wl := (val & -65536) >> 16;
	wr := val & 65535;
	
	FOR i IN 0..11 LOOP
	g1 := (wl>>8) & 255;
	g2 := wl & 255;
	g3 := get_byte(ftable, g2 # get_byte(cr_key, (4*k)%10)) # g1;
	g4 := get_byte(ftable, g3 # get_byte(cr_key, (4*k+1)%10)) # g2;
	g5 := get_byte(ftable, g4 # get_byte(cr_key, (4*k+2)%10)) # g3;
	g6 := get_byte(ftable, g5 # get_byte(cr_key, (4*k+3)%10)) # g4;
	wr := wr # (((g5<<8) + g6) # k);
	k := k + kstep;
	
	g1 := (wr>>8) & 255;
	g2 := wr & 255;
	g3 := get_byte(ftable, g2 # get_byte(cr_key, (4*k)%10)) # g1;
	g4 := get_byte(ftable, g3 # get_byte(cr_key, (4*k+1)%10)) # g2;
	g5 := get_byte(ftable, g4 # get_byte(cr_key, (4*k+2)%10)) # g3;
	g6 := get_byte(ftable, g5 # get_byte(cr_key, (4*k+3)%10)) # g4;
	wl := wl # (((g5<<8) + g6) # k);
	k := k + kstep;
	END LOOP;
	
	RETURN (wr << 16) | (wl & 65535);
	
END
$$ immutable strict language plpgsql;

`)
}

// SetupTriggers is adding update and insert trigger for tables
func SetupTriggers() {
	DB.Exec(`
CREATE OR REPLACE FUNCTION generate_short_id() RETURNS TRIGGER AS $$
BEGIN
	NEW.join_secret := id_to_alpha(skip32(NEW.id, 'sdlkskljsk', true)::int);
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS games_join_secret ON games;
CREATE TRIGGER things_trigger BEFORE INSERT ON games 
FOR EACH ROW EXECUTE PROCEDURE generate_short_id();
`)
}
