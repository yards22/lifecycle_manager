import { Button, Modal } from "@mantine/core";
import { useState } from "react";
import { Plus } from "react-feather";
import AddPollModal from "./AddPollModal";

export default function AddPoll(){
    const [addPollModalOpened , setAddPollModalOpen] = useState(false)
    return (
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
    )
}