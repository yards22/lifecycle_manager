import { Observer } from "mobx-react-lite";
import  { useEffect } from "react";
import { useStores } from "./Logic/Providers/StateProvider";
import FeedBackScreenIndex from "./Screens/FeedbackScreen/Index";
import PollsScreenIndex from "./Screens/PollsScreen/Index";
import StoriesScreenIndex from "./Screens/StoriesScreen/Index";
import styled from "styled-components"
import { BrowserRouter as Router, Routes, Route, Navigate } from "react-router-dom";
import TopBar from "./Organs/Navbar/TopBar/TopBar";
import BottomBar from "./Organs/Navbar/BottomBar";
import LoginScreenIndex from "./Screens/LoginScreen/Index";
import ProtectedRoutes from "./ProtectedRoutes";


const SApp = styled.section`
  width: 100%;
  overflow: auto;
  ::-webkit-scrollbar {
    display: none;
  }
`;

const TOKEN_KEY = "token";

function App() {
  const store = useStores();

  useEffect(()=>{
    handleScreenSizeChange();
    window.addEventListener('resize',handleScreenSizeChange)
    const x = window.localStorage.getItem(TOKEN_KEY)
    if(x){
      store.authStore.SetIsUserLoggedIn(true)
    }else{
      store.authStore.SetIsUserLoggedIn(false)
    }
  },[])

  const handleScreenSizeChange = ()=>{
    store.appStore.setDeviceWidth(window.innerWidth);
      if (window.innerWidth <= 700) {
        store.appStore.setIsPhone(true);
      } else if (window.innerWidth <= 1250) {
        store.appStore.setIsTablet(true);
      } else {
        store.appStore.setIsDesktop(true);
      }
   
  }

  return (
    <Observer>
      {
        ()=>{
          const {appStore,authStore} = store;
          return(
            <Router>
              <SApp>
              { authStore.isUserLoggedIn && appStore.isPhone && <BottomBar/>}
              { authStore.isUserLoggedIn && <TopBar/>}
                 <Routes>
                  <Route path="/login" element={<LoginScreenIndex/>}/>
                      <Route element={<ProtectedRoutes/>}>
                        <Route path="/feedback" element={<FeedBackScreenIndex/>}/>
                        <Route path="/polls" element={<PollsScreenIndex/>}/>
                        <Route path="/stories" element={<StoriesScreenIndex/>}/>
                        <Route path="*" element={<Navigate to={"/feedback"}/>} />
                      </Route>
                 </Routes>
              </SApp>
            </Router>
          )
        }
      }
    </Observer>
  )
  
}

export default App;
