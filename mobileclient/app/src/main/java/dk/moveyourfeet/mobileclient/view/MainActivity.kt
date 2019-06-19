package dk.moveyourfeet.mobileclient.view

import android.content.Intent
import android.os.Bundle
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import android.widget.Toast
import androidx.recyclerview.widget.LinearLayoutManager
import androidx.recyclerview.widget.RecyclerView
import com.marcinmoskala.kotlinandroidviewbindings.bindToSwipeRefresh
import com.marcinmoskala.kotlinandroidviewbindings.bindToVisibility
import dk.moveyourfeet.mobileclient.R
import dk.moveyourfeet.mobileclient.core.data.Game
import dk.moveyourfeet.mobileclient.core.presentation.gamelist.GameListPresenter
import dk.moveyourfeet.mobileclient.core.presentation.gamelist.GameListView
import dk.moveyourfeet.mobileclient.core.repository.GameRepository
import kotlinx.android.synthetic.main.activity_main.*
import kotlinx.android.synthetic.main.content_main.*

class MainActivity : BaseActivity(), GameListView {

  private val gameRepository by GameRepository.lazyGet()

  private val gameListPresenter by presenter { GameListPresenter(this, gameRepository) }

  override var loading by bindToVisibility(R.id.progressBar)
  override var refresh by bindToSwipeRefresh(R.id.swiperefresh)

  override fun onCreate(savedInstanceState: Bundle?) {
    setContentView(R.layout.activity_main)
    super.onCreate(savedInstanceState)
    setSupportActionBar(toolbar)
    swiperefresh.setOnRefreshListener { gameListPresenter.onRefresh() }
    gameListView.layoutManager = LinearLayoutManager(this)
  }

  fun createGame(view: View) {
    val intent = Intent(this, CreateGameActivity::class.java)
    startActivity(intent)
  }

  override fun showGameList(games: List<Game>) {
    gameListView.adapter = GameAdapter(games)
  }

  inner class GameAdapter(private val games: List<Game>) : RecyclerView.Adapter<GameAdapter.GameViewHolder>() {

    override fun onCreateViewHolder(parent: ViewGroup, viewType: Int): GameViewHolder {
      val view = LayoutInflater.from(parent.context).inflate(R.layout.game_row_item, parent, false)
      return GameViewHolder(view)
    }

    override fun getItemCount() = games.size

    override fun onBindViewHolder(holder: GameViewHolder, position: Int) {
      holder.textView.text = games[position].name
    }

    inner class GameViewHolder(val view: View) : RecyclerView.ViewHolder(view) {
      val textView: TextView

      init {
        view.setOnClickListener {
          Toast.makeText(
            this@MainActivity,
            "Game '" + games[adapterPosition].name + "' clicked!",
            Toast.LENGTH_SHORT
          ).show()
        }
        textView = view.findViewById(R.id.textView)
      }
    }
  }
}
