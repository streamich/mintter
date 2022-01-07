use crate::{Error, Hook, Host, ModuleLoader};
use async_trait::async_trait;
use common::{MenuItem, MenuKind};
use tracing::instrument;

#[async_trait]
pub trait HostExt {
  async fn menu(&self, kind: &MenuKind) -> Result<Vec<MenuItem>, Error>;
}

#[async_trait]
impl<L: ModuleLoader + Sync> HostExt for Host<L> {
  #[instrument(level = "trace")]
  async fn menu(&self, kind: &MenuKind) -> Result<Vec<MenuItem>, Error> {
    let menu_items = self
      .hook_parallel::<_, Vec<MenuItem>>(Hook::Menu, kind)
      .await?
      .into_iter()
      .flatten() // each plugin can return an array of menu items
      .collect();

    Ok(menu_items)
  }
}