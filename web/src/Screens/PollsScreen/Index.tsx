import { Button, Modal } from '@mantine/core';
import { Observer } from 'mobx-react-lite'
import { useEffect, useState } from 'react'
import styled from 'styled-components';
import PollIndex from './PollCards';
import { Plus } from 'react-feather';
import AddPollModal from './AddPollModal';
import { useStores } from '../../Logic/Providers/StateProvider';
import AddPoll from './AddPoll';
import PollCards from './PollCards';

const SPollIndex = styled.div`
  display: flex;
  flex-direction: column;
  width: 100%;
`
const SPollWrapContainer = styled.div`
  display: grid;
  grid-template-columns: ${(p)=> p.theme.deviceWidth > 1600 ? "auto auto auto": p.theme.deviceWidth < 800 ? "auto":"auto auto"};
  width: 100%;
`

function PollsScreenIndex() {
  const stores = useStores();
  return (
    <Observer>
       {
         () =>{
           const {appStore,pollsStore} = stores
           if (pollsStore.pollsArray?.length){
             return(
               <SPollIndex>
                  <AddPoll/>
                  <SPollWrapContainer theme={{deviceWidth : appStore.deviceWidth}}>
                     {
                      pollsStore.pollsArray.map((item,index)=>
                         <PollCards polls={item} key={`normal_post_${index}`}/>
                      )
                     }
                  </SPollWrapContainer>
               </SPollIndex>
             )
           }
           return (
             <p>No Polls</p>
           )
         }
       }
    </Observer>
  )
}

export default PollsScreenIndex