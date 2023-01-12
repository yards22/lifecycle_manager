import { MantineProvider } from '@mantine/core';
import { Observer } from 'mobx-react-lite';
import React from 'react'
import styled from 'styled-components';
import { StoresContext } from './Logic/Providers/StateProvider';
import AppStore from './Logic/State/AppStore'

const SProvidedApp = styled.div`
  display: flex;
  justify-content: center;
`
interface ProvidedAppProps {
    children?: React.ReactNode;
  }

function ProvidedApp(props:ProvidedAppProps) {
  
  const appStore = new AppStore();
  return (
    <SProvidedApp>
      <StoresContext.Provider
        value={{
          appStore
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