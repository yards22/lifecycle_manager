import { Button, Modal } from '@mantine/core';
import { Observer } from 'mobx-react-lite'
import { useState } from 'react'
import styled from 'styled-components';
import PollIndex from './Poll/Index';
import { Plus } from 'react-feather';
import AddPollModal from './AddPollModal';

const SPollIndex = styled.div`
  display: flex;
  flex-wrap: wrap;
  justify-content: space-around;
`

function PollsScreenIndex() {
  const [addPollModalOpened , setAddPollModalOpen] = useState(false)
  return (
    <Observer>
       {
         () =>{
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
                <PollIndex/>
                <PollIndex/>
                <PollIndex/>
                <PollIndex/>
                <PollIndex/>
             </SPollIndex>
           )
         }
       }
    </Observer>
  )
}

export default PollsScreenIndex