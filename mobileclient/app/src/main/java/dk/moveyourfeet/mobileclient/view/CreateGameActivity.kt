package dk.moveyourfeet.mobileclient.view

import android.content.Intent
import android.os.Bundle
import android.view.View
import dk.moveyourfeet.mobileclient.R
import dk.moveyourfeet.mobileclient.core.presentation.creategame.CreateGameView
import dk.moveyourfeet.mobileclient.core.repository.GameRepository

class CreateGameActivity : BaseActivity(), CreateGameView {

  private val gameRepository by GameRepository.lazyGet()

  override fun onCreate(savedInstanceState: Bundle?) {
    setContentView(R.layout.activity_create_game)
    super.onCreate(savedInstanceState)
  }

  fun startGame(view: View) {
    val intent = Intent(this, MapActivity::class.java)
    startActivity(intent)
  }
}
