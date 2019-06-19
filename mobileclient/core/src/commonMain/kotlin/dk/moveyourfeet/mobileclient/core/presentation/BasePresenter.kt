package dk.moveyourfeet.mobileclient.core.presentation

import kotlinx.coroutines.Job

abstract class BasePresenter : Presenter {

  protected var jobs : List<Job> = emptyList()

  override fun onDestroy() {
    jobs.filter { !it.isCancelled }.forEach { it.cancel() }
  }
}
