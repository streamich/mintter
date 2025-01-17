use tauri::{
  api::shell::open, window::WindowBuilder, AppHandle, Manager, Runtime, Window, WindowUrl,
};

#[cfg(target_os = "macos")]
use crate::window::{close_all_windows, new_window};
#[cfg(target_os = "macos")]
use anyhow::bail;
#[cfg(target_os = "macos")]
use log::error;
#[cfg(target_os = "macos")]
use tauri::{CustomMenuItem, Menu, MenuItem, Submenu, WindowMenuEvent};

#[tauri::command(async)]
pub fn open_about<R: Runtime>(app_handle: AppHandle<R>, window: Window<R>) {
  let package_info = app_handle.package_info();
  let message = format!(
    r#"
    {}

    Version: {}
    Commit: {}

    Copyright © 2019-2022 {}.
    Some rights reserved.
  "#,
    package_info.description,
    package_info.version,
    std::option_env!("GITHUB_SHA").unwrap_or("N/A"),
    package_info.authors,
  );

  tauri::api::dialog::message(Some(&window), &package_info.name, message);
}

#[tauri::command(async)]
pub fn open_preferences<R: Runtime>(app_handle: AppHandle<R>) -> tauri::Result<()> {
  if let Some(window) = app_handle.get_window("preferences") {
    window.set_focus()?;
  } else {
    let win = WindowBuilder::new(
      &app_handle,
      "preferences",
      WindowUrl::App("/settings".into()),
    )
    .title("Preferences");

    #[cfg(not(target_os = "macos"))]
    let win = { win.decorations(false) };

    win.build()?;
  }

  Ok(())
}

#[tauri::command(async)]
pub fn open_documentation<R: Runtime>(app_handle: AppHandle<R>) {
  open(&app_handle.shell_scope(), "https://mintter.com", None).unwrap();
}

#[tauri::command(async)]
pub fn open_release_notes<R: Runtime>(app_handle: AppHandle<R>) {
  open(&app_handle.shell_scope(), "https://mintter.com", None).unwrap();
}

#[tauri::command(async)]
pub fn open_acknowledgements<R: Runtime>(_app_handle: AppHandle<R>) {
  todo!()
}

#[cfg(target_os = "macos")]
pub fn get_menu() -> Menu {
  let app_menu = Menu::new()
    .add_item(CustomMenuItem::new("about", "About Mintter"))
    .add_native_item(MenuItem::Separator)
    .add_item(CustomMenuItem::new("preferences", "Preferences...").accelerator("CmdOrControl+,"))
    .add_native_item(MenuItem::Separator)
    .add_native_item(MenuItem::Hide)
    .add_native_item(MenuItem::HideOthers)
    .add_native_item(MenuItem::Separator)
    .add_native_item(MenuItem::Quit);

  let file_menu = Menu::new()
    .add_item(CustomMenuItem::new("new_window", "New Window").accelerator("CmdOrControl+N"))
    .add_native_item(MenuItem::Separator)
    .add_native_item(MenuItem::CloseWindow)
    .add_item(
      CustomMenuItem::new("close_all_windows", "Close All Windows")
        .accelerator("Alt+Shift+CmdOrControl+W"),
    );

  let edit_menu = Menu::new()
    .add_native_item(MenuItem::Undo)
    .add_native_item(MenuItem::Redo)
    .add_native_item(MenuItem::Separator)
    .add_native_item(MenuItem::Cut)
    .add_native_item(MenuItem::Copy)
    .add_native_item(MenuItem::Paste)
    .add_item(CustomMenuItem::new("select_all", "Select All").accelerator("CmdOrControl+A"))
    .add_item(CustomMenuItem::new("find", "Find...").accelerator("CmdOrControl+F"));

  let format_menu = Menu::new()
    .add_item(CustomMenuItem::new("strong", "Strong").accelerator("CmdOrControl+B"))
    .add_item(CustomMenuItem::new("emphasis", "Emphasis").accelerator("CmdOrControl+I"))
    .add_item(CustomMenuItem::new("code", "Code").accelerator("CmdOrControl+E"))
    .add_item(CustomMenuItem::new("underline", "Underline").accelerator("CmdOrControl+U"))
    .add_item(CustomMenuItem::new("strikethrough", "Strikethrough"))
    .add_item(CustomMenuItem::new("subscript", "Subscript"))
    .add_item(CustomMenuItem::new("superscript", "Superscript"))
    .add_native_item(MenuItem::Separator)
    .add_item(CustomMenuItem::new("heading", "Heading").accelerator("Shift+CmdOrControl+H"))
    .add_item(CustomMenuItem::new("statement", "Statement").accelerator("Shift+CmdOrControl+S"))
    .add_item(CustomMenuItem::new("blockquote", "Blockquote").accelerator("Shift+CmdOrControl+Q"))
    .add_item(CustomMenuItem::new("codeblock", "Code Block").accelerator("Shift+CmdOrControl+E"))
    .add_native_item(MenuItem::Separator)
    .add_item(
      CustomMenuItem::new("unordered_list", "Bullet List").accelerator("Shift+CmdOrControl+7"),
    )
    .add_item(
      CustomMenuItem::new("ordered_list", "Numbered List").accelerator("Shift+CmdOrControl+8"),
    )
    .add_item(CustomMenuItem::new("group", "Plain List").accelerator("Shift+CmdOrControl+9"));

  let view_menu = Menu::new()
    .add_item(CustomMenuItem::new("reload", "Reload").accelerator("CmdOrControl+R"))
    .add_item(
      CustomMenuItem::new("quick_switcher", "Quick Switcher...").accelerator("CmdOrControl+K"),
    );

  let help_menu = Menu::new()
    .add_item(CustomMenuItem::new("documentation", "Documentation"))
    .add_item(CustomMenuItem::new("release_notes", "Release Notes"))
    .add_item(CustomMenuItem::new("acknowledgements", "Acknowledgements"));

  Menu::new()
    .add_submenu(Submenu::new("Mintter", app_menu))
    .add_submenu(Submenu::new("File", file_menu))
    .add_submenu(Submenu::new("Edit", edit_menu))
    .add_submenu(Submenu::new("Format", format_menu))
    .add_submenu(Submenu::new("View", view_menu))
    .add_submenu(Submenu::new("Help", help_menu))
}

#[cfg(target_os = "macos")]
pub fn event_handler(event: WindowMenuEvent) {
  if let Err(err) = event_handler_inner(event) {
    error!("Failed to handle menu event {}", err);
  }
}

#[cfg(target_os = "macos")]
pub fn event_handler_inner(event: WindowMenuEvent) -> anyhow::Result<()> {
  match event.menu_item_id() {
    "new_window" => {
      new_window(event.window().clone())?;
    }
    "reload" => {
      event.window().eval("location.reload()")?;
    }
    "close_all_windows" => {
      close_all_windows(event.window().app_handle())?;
    }
    "documentation" => {
      open_documentation(event.window().app_handle());
    }
    "release_notes" => {
      open_release_notes(event.window().app_handle());
    }
    "acknowledgements" => {
      open_acknowledgements(event.window().app_handle());
    }
    "preferences" => {
      open_preferences(event.window().app_handle())?;
    }
    "about" => {
      open_about(event.window().app_handle(), event.window().clone());
    }
    "find" => {
      event.window().emit("open_find", ())?;
    }
    "quick_switcher" => {
      event.window().emit("open_quick_switcher", ())?;
    }
    "select_all" => {
      event.window().emit("select_all", ())?;
    }
    "strong" | "emphasis" | "code" | "underline" | "strikethrough" | "subscript"
    | "superscript" => {
      event.window().emit("format_mark", event.menu_item_id())?;
    }
    "heading" | "statement" | "blockquote" | "codeblock" => {
      event.window().emit("format_block", event.menu_item_id())?;
    }
    "unordered_list" | "ordered_list" | "group" => {
      event.window().emit("format_list", event.menu_item_id())?;
    }
    id => bail!("Unhandled menu item \"{}\"", id),
  }

  Ok(())
}
