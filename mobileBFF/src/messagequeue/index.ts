
/* eslint-disable no-unused-vars */
import { connect, NatsConnectionOptions, Payload, Client } from 'ts-nats';

// // ...
// try {
//     let nc = await connect({servers: ['nats://localhost:4222', 'tls://localhost:4443']});
//     // Do something with the connection
// } catch(ex) {
//     // handle the error
// }

// let msg = await nc.request('time', 1000);
// t.log('the time is', msg.data);
// nc.close();

import { createServer, Server } from 'http';
import socketIo from 'socket.io';

import { CreateGame, Message } from '../model';

export class MQ {
    // Default value
    public static readonly NATS_URI: string = "nats://nats:4222";
    private nc: Client;
    private server: Server;
    private io: SocketIO.Server;
    private natsUri: string;

    constructor() {
        this.init()
    }

    async init() {
        this.config()
        this.nc = await this.connectToNats() 
    }

    public async createGame(createGame: CreateGame): Promise<string> {
        let msg = await this.nc.request('time', 1000);
        return msg.data;
    }

    private async connectToNats(): Promise<Client> {
        try {
            let nc = await connect({'maxReconnectAttempts': -1, 'reconnectTimeWait': 250, servers: [this.natsUri] });
            return nc
            // Do something with the connection
        } catch (ex) {
            console.log('[ERROR]: Could not connect to NATS..');
            throw new Error(ex);
        }
    }

    private config(): void {
        this.natsUri = process.env.NATS_URI || MQ.NATS_URI;
    }
}
