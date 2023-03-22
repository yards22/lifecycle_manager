import { Burger, Title, useMantineTheme } from '@mantine/core';
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
    const [isLoggedIn, setIsLoggedIn] = useState(false)
    
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
                justifyContent: "space-between",
                alignItems: "center",
            }}
            >
            <Burger 
                opened={isNavBarOpened}
                color = {'white'}
                onClick = {()=>{setIsNavBarOpened(!isNavBarOpened)}}
            />
            <Title color={"white"} order={5}>
                Admin 22 Yards 
            </Title>
            <LogOut
                color = {'white'}
                onClick = {()=>{
                    console.log("h")
                    stores.authStore.PerformLogout().then(
                        ()=>{
                            navigate( "/login")
                        })
                        .catch(e=>{
                            throw e
                        })
                }}
             />
            { isNavBarOpened && <NavBarMobile setIsNavBarOpened={setIsNavBarOpened}/> }
            </div>
        )
        }}
        </Observer>
    )
}

export default TopBarMobile