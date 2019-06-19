package dk.moveyourfeet.mobileclient.core.presentation

interface BaseView {
  fun logError(error: Throwable)
  fun showError(error: Throwable)
}
