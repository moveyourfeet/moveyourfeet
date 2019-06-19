package dk.moveyourfeet.mobileclient

import android.app.Application

class App : Application() {

  override fun onCreate() {
    super.onCreate()
  }

  companion object {
    var baseUrl: String? = null
  }
}
