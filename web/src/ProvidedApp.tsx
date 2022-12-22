import { MantineProvider } from '@mantine/core';
import React from 'react'
import { StoresContext } from './Logic/Providers/StateProvider';
import AppStore from './Logic/State/AppStore'

interface ProvidedAppProps {
    children?: React.ReactNode;
  }

function ProvidedApp(props:ProvidedAppProps) {
  
  const appStore = new AppStore();
  return (
    <StoresContext.Provider
       value={{
        appStore
       }}
    >
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
    </StoresContext.Provider>
  )
}

export default ProvidedApp