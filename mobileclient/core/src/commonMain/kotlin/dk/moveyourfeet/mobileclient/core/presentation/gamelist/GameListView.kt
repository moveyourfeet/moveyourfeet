package dk.moveyourfeet.mobileclient.core.presentation.gamelist

import dk.moveyourfeet.mobileclient.core.data.Game
import dk.moveyourfeet.mobileclient.core.presentation.BaseView

interface GameListView: BaseView {
  var loading: Boolean
  var refresh: Boolean
  fun showGameList(games : List<Game>)
}
