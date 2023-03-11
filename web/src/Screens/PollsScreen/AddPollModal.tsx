import { Button, Input, Textarea } from '@mantine/core'
import { useState } from 'react'
import {  Trash2 } from 'react-feather';
import styled from 'styled-components'
import { useStores } from '../../Logic/Providers/StateProvider';


const SAddPollModal = styled.div`
    display: flex;
    flex-direction: column;
    width: 100%;
`

function AddPollModal() {
  const [pollOptions , setPollOptions] = useState<string[]>([""]);
  const [pollQuestion, setPollQuestion] = useState("")
  const {pollsStore,authStore} = useStores();

  function handleOnChangeOfOptions(option:string,index:number){
     let w = pollOptions;
     w[index] = option;
     setPollOptions([...w])
  }

  function handlePollQuestionChange(e:any){
    setPollQuestion(e.target.value)
  }

  function handleDeleteOption(removeIndex:number){
    console.log(removeIndex)
    pollOptions.splice(removeIndex,1)
    setPollOptions([...pollOptions])
  }

  function handleSubmitThePoll(){
    console.log("Question",pollQuestion);
    console.log("options",pollOptions)
    pollsStore.AddPoll(pollQuestion,pollOptions, authStore.token?authStore.token:"")
  }

  return (
     <SAddPollModal>
        <label>
            Enter the Poll Question
            <Textarea mb={6} onChange={handlePollQuestionChange}/>
        </label>
        {
            pollOptions.map((each,index)=>{
                return(
                   <label key={index}>
                      Option {index+1}
                       <div style={{
                          display : "flex",
                          alignItems : "center",
                       }}> 
                          <Input 
                            mr={8} 
                            style={{width : "100%"}} 
                            value={each} 
                            onChange={(e:any)=>handleOnChangeOfOptions(e.target.value,index)}
                           />
                          <Trash2 size={25} color={"red"} onClick = {()=>handleDeleteOption(index)}/>
                       </div>
                   </label>
                )
            })
        }
        <div style={{
            display : "flex",
            justifyContent : "space-between",
            marginTop : '20px'
        }}>
            <Button disabled={pollOptions.length>3} variant='outline' onClick={()=>setPollOptions([...pollOptions,""])}>Add Option</Button>
            <Button onClick={handleSubmitThePoll}>Add Poll</Button>
        </div>
     </SAddPollModal>
  )
}

export default AddPollModal