import { MantineProvider } from '@mantine/core';
import { Observer } from 'mobx-react-lite';
import React from 'react'
import styled from 'styled-components';
import { StoresContext } from './Logic/Providers/StateProvider';
import { AuthRepo } from './Logic/Repository/AuthRepo';
import { FeedBackRepo } from './Logic/Repository/FeedBackRepo';
import AppStore from './Logic/State/AppStore'
import { AuthStore } from './Logic/State/AuthStore';
import { FeedBackStore } from './Logic/State/FeedBackStore';
import { Request } from './Logic/Utils/Fetch';

const SProvidedApp = styled.div`
  display: flex;
  justify-content: center;
`
interface ProvidedAppProps {
    children?: React.ReactNode;
  }

const BASE_URL = "http://localhost:3001"

function ProvidedApp(props:ProvidedAppProps) {
  const rq = new Request({});
  const appStore = new AppStore();
  const authStore = new AuthStore(new AuthRepo(BASE_URL,rq))
  const feedBackStore = new FeedBackStore(new FeedBackRepo(BASE_URL+"/feedback",rq))

  return (
    <SProvidedApp>
      <StoresContext.Provider
        value={{
          appStore,
          feedBackStore,
          authStore
        }}
      >
        { 
          <Observer>
          { () => {
            return(
                <MantineProvider
                    withGlobalStyles
                    withNormalizeCSS
                    theme={{
                    loader: "dots",
                    colors: {},
                    }}
                >
                    {props.children}
                </MantineProvider>
            )
          }}
          </Observer>
        }
      </StoresContext.Provider>
    </SProvidedApp>
  )
}

export default ProvidedApp