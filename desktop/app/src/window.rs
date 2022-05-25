use std::{
  collections::HashMap,
  path::PathBuf,
  sync::{Arc, Mutex},
  time::{SystemTime, UNIX_EPOCH},
};

use tauri::{
  plugin::{Builder as PluginBuilder, TauriPlugin},
  AppHandle, Manager, RunEvent, Runtime, State, Window, WindowBuilder, WindowEvent, WindowUrl,
};

#[derive(Debug, Default)]
struct WindowTable(Arc<Mutex<HashMap<String, PathBuf>>>);

#[tauri::command]
async fn update_window_table<R: Runtime>(
  window: Window<R>,
  window_table: State<'_, WindowTable>,
  url: PathBuf,
) -> Result<(), String> {
  let mut window_table = window_table.0.lock().unwrap();

  window_table.insert(window.label().to_string(), url);

  Ok(())
}

#[tauri::command]
async fn open_in_new_window<R: Runtime>(
  app: AppHandle<R>,
  window_table: State<'_, WindowTable>,
  url: PathBuf,
) -> Result<(), String> {
  let mut window_table = window_table.0.lock().unwrap();

  for (label, path) in window_table.iter() {
    if path == &url {
      let win = app.get_window(label).unwrap();

      win.set_focus().map_err(|err| err.to_string())?;
      return Ok(());
    }
  }

  let id = SystemTime::now()
    .duration_since(UNIX_EPOCH)
    .expect("Failed to construct unix timestamp")
    .as_millis()
    .to_string();

  WindowBuilder::new(&app, id.clone(), WindowUrl::App(url.clone()))
    .build()
    .map_err(|err| err.to_string())?;

  window_table.insert(id, url);

  Ok(())
}

pub fn new_window<R: Runtime, M: Manager<R>>(manager: &M) -> tauri::Result<()> {
  let id = SystemTime::now()
    .duration_since(UNIX_EPOCH)
    .expect("Failed to construct unix timestamp")
    .as_millis()
    .to_string();

  WindowBuilder::new(manager, id, WindowUrl::App("index.html".into())).build()?;

  Ok(())
}

pub fn close_all_windows<R: Runtime, M: Manager<R>>(manager: &M) -> tauri::Result<()> {
  for window in manager.windows().values() {
    window.close()?;
  }

  Ok(())
}

pub fn init<R: Runtime>() -> TauriPlugin<R> {
  PluginBuilder::new("window")
    .invoke_handler(tauri::generate_handler![
      update_window_table,
      open_in_new_window
    ])
    .setup(|app_handle| {
      app_handle.manage(WindowTable::default());

      Ok(())
    })
    .on_event(|app_handle, event| {
      if let RunEvent::WindowEvent { label, event, .. } = event {
        if let WindowEvent::Destroyed = event {
          let window_table = app_handle.state::<WindowTable>();

          let mut window_table = window_table.0.lock().unwrap();

          window_table.remove(label);
        }
      }
    })
    .build()
}