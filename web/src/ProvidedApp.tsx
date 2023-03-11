import { MantineProvider } from '@mantine/core';
import { Observer } from 'mobx-react-lite';
import React from 'react'
import styled from 'styled-components';
import { StoresContext } from './Logic/Providers/StateProvider';
import { AuthRepo } from './Logic/Repository/AuthRepo';
import { FeedBackRepo } from './Logic/Repository/FeedBackRepo';
import { PollsRepo } from './Logic/Repository/PollsRepo';
import AppStore from './Logic/State/AppStore'
import { AuthStore } from './Logic/State/AuthStore';
import { FeedBackStore } from './Logic/State/FeedBackStore';
import { PollsStore } from './Logic/State/PollsStore';
import { Request } from './Logic/Utils/Fetch';

const SProvidedApp = styled.div`
  display: flex;
  justify-content: center;
`
interface ProvidedAppProps {
    children?: React.ReactNode;
  }

const BASE_URL = "http://localhost:5000";
const BASE_URL_FOR_IMAGES =
  "https://22yards-image-bucket.s3.ap-south-1.amazonaws.com/";

function ProvidedApp(props:ProvidedAppProps) {
  const rq = new Request({});
  const appStore = new AppStore();
  const authStore = new AuthStore(new AuthRepo(BASE_URL,rq))
  const feedBackStore = new FeedBackStore(new FeedBackRepo(BASE_URL+"/feedback",BASE_URL_FOR_IMAGES,rq))
  const pollsStore = new PollsStore(new PollsRepo(BASE_URL+"/poll",rq))

  return (
    <SProvidedApp>
      <StoresContext.Provider
        value={{
          appStore,
          feedBackStore,
          authStore,
          pollsStore
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