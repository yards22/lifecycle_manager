import { useStores } from "../../../Logic/Providers/StateProvider";
import TopBarDesktop from "./TopBarDesktop";
import TopBarMobile from "./TopBarMobile";
import TopBarTablet from "./TopBarTablet";


function TopBar() {
  const stores = useStores();
  if (stores.appStore.isPhone) {
    return (
      <TopBarMobile/>
    );
  }
  if (stores.appStore.isTablet){
    return (
      <TopBarTablet/>
    )
  }
  return (
    <TopBarDesktop/>
  );
}

export default TopBar;
