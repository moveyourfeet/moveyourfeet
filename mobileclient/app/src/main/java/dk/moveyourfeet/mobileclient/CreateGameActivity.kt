package dk.moveyourfeet.mobileclient

import android.content.Intent
import android.os.Bundle
import android.view.View
import androidx.appcompat.app.AppCompatActivity

class CreateGameActivity : AppCompatActivity() {

  override fun onCreate(savedInstanceState: Bundle?) {
    super.onCreate(savedInstanceState)
    setContentView(R.layout.activity_create_game)
  }

  fun startGame(view: View) {
    val intent = Intent(this, MapActivity::class.java)
    startActivity(intent)
  }
}
