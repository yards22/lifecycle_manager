import { AppShell, Burger,  Header, MediaQuery, Navbar, Text, useMantineTheme } from "@mantine/core";
import { Observer } from "mobx-react-lite";
import React, { useEffect, useRef, useState } from "react";
import { useStores } from "./Logic/Providers/StateProvider";
import FeedBackScreenIndex from "./Screens/FeedbackScreen/Index";
import PollsScreenIndex from "./Screens/PollsScreen/Index";
import StoriesScreenIndex from "./Screens/StoriesScreen/Index";
import styled from "styled-components"
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import TopBar from "./Organs/Navbar/TopBar/TopBar";
import BottomBar from "./Organs/Navbar/BottomBar";
import LoginScreenIndex from "./Screens/LoginScreen/Index";

const STabs = styled.div`
  height: 50px;
  width: 100%;
  border-radius: 5px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
`

const SApp = styled.section`
  max-width: 100%;
  overflow: auto;
  ::-webkit-scrollbar {
    display: none;
  }
`;


function App() {
  const [opened, setOpened] = useState(false);
  const theme = useMantineTheme();
  const store = useStores();
  const divRef:any = useRef(null);

  useEffect(()=>{
    handleScreenSizeChange();
    window.addEventListener('resize',handleScreenSizeChange)
  },[])

  const handleScreenSizeChange = ()=>{
    //  store.appStore.SetCurrentWindowSize(divRef.current.clientWidth)
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
              { authStore.isUserLoggedIn && appStore.isPhone && <BottomBar/>}
              { authStore.isUserLoggedIn && <TopBar/>}
              <SApp>
                 <Routes>
                   <Route path="/" element={<LoginScreenIndex/>} />
                   <Route path="/feedback" element={<FeedBackScreenIndex/>}/>
                   <Route path="/polls" element={<PollsScreenIndex/>}/>
                   <Route path="/stories" element={<StoriesScreenIndex/>}/>
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
