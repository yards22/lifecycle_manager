import { Button, Modal } from '@mantine/core';
import { Observer } from 'mobx-react-lite'
import { useEffect, useState } from 'react'
import styled from 'styled-components';
import PollIndex from './Poll/Index';
import { Plus } from 'react-feather';
import AddPollModal from './AddPollModal';
import { useStores } from '../../Logic/Providers/StateProvider';

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
  const [addPollModalOpened , setAddPollModalOpen] = useState(false)
  const stores = useStores();
  useEffect(()=>{
    stores.pollsStore.GetPolls(stores.authStore.token ? stores.authStore.token:"")
  },[])
  return (
    <Observer>
       {
         () =>{
           const {appStore} = stores
           return(
             <SPollIndex>
                <div style={{
                    width : "100%",
                    display : "flex",
                    justifyContent : "flex-end"
                  }}>
                      <Button
                        mb={20}
                        onClick={() => setAddPollModalOpen(true)}
                        rightIcon = {<Plus size={20}/>}
                      >
                        ADD POLL
                      </Button>
                      <Modal
                        opened={addPollModalOpened}
                        onClose={() => setAddPollModalOpen(false)}
                        title="Add A New Poll"
                      >
                        <AddPollModal/>
                      </Modal>
                </div>
                <SPollWrapContainer theme={{deviceWidth : appStore.deviceWidth}}>
                  <PollIndex/>
                  <PollIndex/>
                  <PollIndex/>
                  <PollIndex/>
                  <PollIndex/>
                </SPollWrapContainer>
             </SPollIndex>
           )
         }
       }
    </Observer>
  )
}

export default PollsScreenIndex