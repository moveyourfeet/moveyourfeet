package dk.moveyourfeet.mobileclient.view

import android.os.Bundle
import android.util.Log
import android.widget.Toast
import androidx.appcompat.app.AppCompatActivity
import dk.moveyourfeet.mobileclient.BuildConfig
import dk.moveyourfeet.mobileclient.R
import dk.moveyourfeet.mobileclient.core.presentation.BasePresenter
import dk.moveyourfeet.mobileclient.core.presentation.BaseView
import java.util.concurrent.CancellationException

abstract class BaseActivity : AppCompatActivity(), BaseView {

  protected fun <T: BasePresenter> presenter(init: () -> T) = lazy(init).also { lazyPresenters += it }

  private var lazyPresenters: List<Lazy<BasePresenter>> = emptyList()

  override fun onCreate(savedInstanceState: Bundle?) {
    super.onCreate(savedInstanceState)
    lazyPresenters.forEach { it.value.onCreate() }
  }

  override fun onDestroy() {
    super.onDestroy()
    lazyPresenters.forEach { it.value.onDestroy() }
  }

  override fun showError(error: Throwable) {
    logError(error)
    when (error) {
      is CancellationException -> { }
      else -> Toast.makeText(this, String.format(getString(R.string.error), error.message), Toast.LENGTH_LONG).show()
    }
  }

  override fun logError(error: Throwable) {
    if (BuildConfig.DEBUG) {
      Log.e(this::class.simpleName, error.message, error)
    }
  }
}
