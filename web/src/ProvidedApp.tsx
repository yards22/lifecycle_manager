import { MantineProvider } from '@mantine/core';
import { Observer } from 'mobx-react-lite';
import React from 'react'
import styled from 'styled-components';
import { StoresContext } from './Logic/Providers/StateProvider';
import { FeedBackRepo } from './Logic/Repository/FeedBackRepo';
import AppStore from './Logic/State/AppStore'
import { FeedBackStore } from './Logic/State/FeedBack';
import { Request } from './Logic/Utils/Fetch';

const SProvidedApp = styled.div`
  display: flex;
  justify-content: center;
`
interface ProvidedAppProps {
    children?: React.ReactNode;
  }

const BASE_URL = "localhost:4000"

function ProvidedApp(props:ProvidedAppProps) {
  const rq = new Request({});
  const appStore = new AppStore();
  const feedBackStore = new FeedBackStore(new FeedBackRepo(BASE_URL,rq))

  return (
    <SProvidedApp>
      <StoresContext.Provider
        value={{
          appStore,
          feedBackStore
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