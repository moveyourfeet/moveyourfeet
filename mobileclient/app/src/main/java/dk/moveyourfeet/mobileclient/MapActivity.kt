package dk.moveyourfeet.mobileclient

import android.os.Bundle
import android.widget.Toast
import androidx.appcompat.app.AppCompatActivity
import androidx.core.content.ContextCompat
import com.google.android.material.snackbar.Snackbar
import com.mapbox.android.core.permissions.PermissionsListener
import com.mapbox.android.core.permissions.PermissionsManager
import com.mapbox.mapboxsdk.Mapbox
import com.mapbox.mapboxsdk.camera.CameraUpdateFactory
import com.mapbox.mapboxsdk.location.LocationComponentOptions
import com.mapbox.mapboxsdk.location.modes.CameraMode
import com.mapbox.mapboxsdk.location.modes.RenderMode
import com.mapbox.mapboxsdk.maps.MapboxMap
import com.mapbox.mapboxsdk.maps.OnMapReadyCallback
import com.mapbox.mapboxsdk.maps.Style
import kotlinx.android.synthetic.main.activity_map.*

const val EXTRA_MESSAGE = "dk.moveyourfeet.mobileclient.MESSAGE"

class MapActivity : AppCompatActivity(), OnMapReadyCallback, PermissionsListener
{
  private var permissionsManager: PermissionsManager = PermissionsManager(this)

  private lateinit var mapboxMap: MapboxMap

  override fun onCreate(savedInstanceState: Bundle?)
  {
    super.onCreate(savedInstanceState)
    Mapbox.getInstance(this, getString(R.string.mapbox_access_token))
    setContentView(R.layout.activity_map)
    mapView.onCreate(savedInstanceState)
    mapView.getMapAsync(this)
  }

  override fun onMapReady(mapboxMap: MapboxMap) {
    this.mapboxMap = mapboxMap
    mapboxMap.setStyle(Style.OUTDOORS) {
      // Map is set up and the style has loaded. Now you can add data or make other map adjustments
      enableLocationComponent(it)
    }
  }

  private fun enableLocationComponent(loadedMapStyle: Style) {
    // Check if permissions are enabled and if not request
    if (PermissionsManager.areLocationPermissionsGranted(this)) {

      val options = LocationComponentOptions.builder(this)
        .trackingGesturesManagement(true)
        .accuracyColor(ContextCompat.getColor(this, R.color.mapboxGreen))
        .build()

      // Get an instance of the component
      val locationComponent = mapboxMap.locationComponent

      // Activate the component
      locationComponent.activateLocationComponent(this, loadedMapStyle)

      // Apply the options to the LocationComponent
      locationComponent.applyStyle(options)

      // Enable to make component visible
      locationComponent.isLocationComponentEnabled = true

      // Set the component's camera mode
      locationComponent.cameraMode = CameraMode.TRACKING
      locationComponent.renderMode = RenderMode.COMPASS

      mapboxMap.animateCamera(CameraUpdateFactory.zoomTo(11.0))

    } else {
      permissionsManager = PermissionsManager(this)
      permissionsManager.requestLocationPermissions(this)
    }
  }

  override fun onRequestPermissionsResult(requestCode: Int, permissions: Array<String>, grantResults: IntArray) {
    permissionsManager.onRequestPermissionsResult(requestCode, permissions, grantResults)
  }

  override fun onExplanationNeeded(permissionsToExplain: List<String>) {
    Snackbar.make(mapView, R.string.user_location_permission_explanation, Snackbar.LENGTH_LONG).show()
  }

  override fun onPermissionResult(granted: Boolean) {
    if (granted) {
      enableLocationComponent(mapboxMap.style!!)
    } else {
      Toast.makeText(this, R.string.user_location_permission_not_granted, Toast.LENGTH_LONG).show()
      finish()
    }
  }

  public override fun onStart() {
    super.onStart()
    mapView.onStart()
  }

  public override fun onResume() {
    super.onResume()
    mapView.onResume()
  }

  public override fun onPause() {
    super.onPause()
    mapView.onPause()
  }

  public override fun onStop() {
    super.onStop()
    mapView.onStop()
  }

  override fun onLowMemory() {
    super.onLowMemory()
    mapView.onLowMemory()
  }

  override fun onDestroy() {
    super.onDestroy()
    mapView.onDestroy()
  }

  override fun onSaveInstanceState(outState: Bundle) {
    super.onSaveInstanceState(outState)
    mapView.onSaveInstanceState(outState)
  }

}
