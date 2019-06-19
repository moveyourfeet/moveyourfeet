package dk.moveyourfeet.mobileclient.core.presentation.gamelist

import dk.moveyourfeet.mobileclient.core.data.Game
import dk.moveyourfeet.mobileclient.core.presentation.BasePresenter
import dk.moveyourfeet.mobileclient.core.repository.GameRepository
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.GlobalScope
import kotlinx.coroutines.launch

class GameListPresenter(
  private val view: GameListView,
  private val repository: GameRepository) : BasePresenter() {

  private var visibleGames: List<Game>? = null

  override fun onCreate() {
    view.loading = true
    refreshList()
  }

  fun onRefresh() {
    view.refresh = true
    refreshList()
  }

  private fun refreshList() {
    jobs += GlobalScope.launch(Dispatchers.Main) {
      try {
        val gamesData = repository.fetchGamesNearby()
        if (gamesData == visibleGames) {
          return@launch
        }
        visibleGames = gamesData

        view.showGameList(gamesData)
      } catch (e: Throwable) {
        view.showError(e)
      } finally {
        view.refresh = false
        view.loading = false
      }
    }
  }

}
