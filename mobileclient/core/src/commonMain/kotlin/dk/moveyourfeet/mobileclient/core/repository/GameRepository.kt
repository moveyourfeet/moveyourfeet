package dk.moveyourfeet.mobileclient.core.repository

import dk.moveyourfeet.mobileclient.core.Provider
import dk.moveyourfeet.mobileclient.core.data.Game
import dk.moveyourfeet.mobileclient.core.data.Player
import io.ktor.client.HttpClient

class GameRepository {

  private val client = HttpClient()

  suspend fun fetchGamesNearby(): List<Game> {
    return listOf(
      Game("Fang mig"),
      Game("Et andet spil"))
  }

  suspend fun createGame(name: String, player: Player): Game {
    TODO()
  }

  suspend fun startGame(game: Game) {
    TODO()
  }

  suspend fun joinGame(game: Game, player: Player) {
    TODO()
  }

  companion object : Provider<GameRepository>() {
    override fun create() = GameRepository()
  }
}
