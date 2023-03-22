import { Burger, Loader, Title, useMantineTheme } from '@mantine/core';
import { Observer } from 'mobx-react-lite';
import React, { useState } from 'react'
import { LogOut, Search } from 'react-feather';
import { useStores } from '../../../Logic/Providers/StateProvider';
import NavBarMobile from '../NavbarMobile/Index';
import { Navigate, useNavigate } from 'react-router-dom';

function TopBarMobile() {
    const stores = useStores();
    const mantineTheme = useMantineTheme();
    const [isNavBarOpened, setIsNavBarOpened] = useState(false)
    const navigate = useNavigate();

    // if(stores.authStore.isLoading)
    // return (
    //     <section
    //     style={{
    //       height: "100vh",
    //       width: "100vh",
    //       display: "flex",
    //       justifyContent: "center",
    //       alignItems: "center"
    //     }}
    //   >
    //     <Loader variant="bars" size={"sm"} />
    //   </section>
    // )
    
    return (
        <Observer>
        {()=>{
        const { appStore } = stores;
        return(
            <div
            style={{
              background: mantineTheme.colors[mantineTheme.primaryColor][6],
              position: "fixed",
              top: "0",
              left: "0",
              zIndex: "100",
              right: "0",
              height: "60px",
              width: "100%",
              display: "flex",
              justifyContent: "center",
              alignItems: "center"
            }}
          >
            <Burger
              opened={isNavBarOpened}
              color={"white"}
              style={{
                position: "absolute",
                left: "13px"
              }}
              onClick={() => {
                setIsNavBarOpened(!isNavBarOpened);
              }}
            />
            <h3 style={{color:"white"}}>22 Yardz Admin</h3>
            <LogOut
                style={{
                    position: "absolute",
                    right: "13px",
                    color: "white"
                }}
                onClick = {()=>{
                    stores.authStore.PerformLogout().then(
                        ()=>{
                            navigate( "/login")
                        })
                        .catch(e=>{
                            throw e
                        })
                }}
            />
            {isNavBarOpened && (
              <NavBarMobile
                setIsNavBarOpened={(x: boolean) => setIsNavBarOpened(x)}
              />
            )}
          </div>
        )
    }}
        </Observer>
    )
}

export default TopBarMobile
