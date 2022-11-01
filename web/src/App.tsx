import { AppShell, Burger,  Header, MediaQuery, Navbar, Text, useMantineTheme } from "@mantine/core";
import { Observer } from "mobx-react-lite";
import React, { useEffect, useRef, useState } from "react";
import { useStores } from "./Logic/Providers/StateProvider";
import FeedBackScreenIndex from "./Screens/FeedbackScreen/Index";
import PollsScreenIndex from "./Screens/PollsScreen/Index";
import StoriesScreenIndex from "./Screens/StoriesScreen/Index";
import styled from "styled-components"

const STabs = styled.div`
  height: 50px;
  width: 100%;
  border-radius: 5px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
`


function App() {
  const [opened, setOpened] = useState(false);
  const theme = useMantineTheme();
  const stores = useStores();
  const divRef:any = useRef(null);

  useEffect(()=>{
    handleScreenSizeChange();
    window.addEventListener('resize',handleScreenSizeChange)
  },[])

  const handleScreenSizeChange = ()=>{
     stores.appStore.SetCurrentWindowSize(divRef.current.clientWidth)
  }

  return (
    <Observer>
      {
        ()=>{
          const {appStore} = stores;
          return(
            <AppShell
                ref={divRef}
                navbarOffsetBreakpoint="sm"
                asideOffsetBreakpoint="sm"
                navbar={
                  <Navbar p="md" hiddenBreakpoint="sm" hidden={!opened} width={{ sm: 200, lg: 300 }}>
                    <STabs 
                       onClick={()=> appStore.SetCurrentTab(0)}
                       style = {{
                          color : appStore.currentTab === 0 ? "white" : "black",
                          backgroundColor : appStore.currentTab === 0 ? theme.colors.blue[6] : "white"
                       }}
                     >Polls Section</STabs>
                    <STabs 
                      onClick={()=> appStore.SetCurrentTab(1)}
                      style = {{
                        color : appStore.currentTab === 1 ? "white" : "black",
                        backgroundColor : appStore.currentTab === 1 ? theme.colors.blue[6] : "white"
                     }}
                    >Feedbacks Section</STabs>
                    <STabs 
                      onClick={()=> appStore.SetCurrentTab(2)}
                      style = {{
                        color : appStore.currentTab === 2 ? "white" : "black",
                        backgroundColor : appStore.currentTab === 2 ? theme.colors.blue[6] : "white"
                     }}
                      >Stories Section</STabs>
                  </Navbar>
                }
                header={
                  <Header height={60} p="md">
                    <div style={{ display: 'flex', alignItems: 'center', height: '100%',  }}>
                      <MediaQuery largerThan="sm" styles={{ display: 'none' }}>
                        <Burger
                          opened={opened}
                          onClick={() => setOpened((o) => !o)}
                          size="sm"
                          color={theme.colors.gray[6]}
                          mr="xl"
                        />
                      </MediaQuery>
                      <Text>ADMIN 22 Yards</Text>
                    </div>
                  </Header>
                }
            > 
              {
                appStore.currentTab === 0 &&  <PollsScreenIndex/>
              }
              {
                appStore.currentTab === 1 && <FeedBackScreenIndex/>
              }
              {
                appStore.currentTab === 2 && <StoriesScreenIndex/>
              }
            </AppShell>
          )
        }
      }
    </Observer>
  )
  
}

export default App;
