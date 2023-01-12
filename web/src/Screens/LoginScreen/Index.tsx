import { Button, Card, Input } from '@mantine/core'
import React, { useState } from 'react'
import styled from 'styled-components'
import { useStores } from '../../Logic/Providers/StateProvider'

const SLoginScreenIndex = styled.div`
    display: flex;
    width: 100%;
    height: 100%;
    justify-content: center;
    align-items: center;

`

function LoginScreenIndex() {
  const stores = useStores();
  const [mailId,setMailId] = useState("")

  function handleMailIdChange(e:any){
    setMailId(e.target.value)
  }

  function handleSendAuthOTP(){
    stores.authStore.SendAuthOTP(mailId)
  }

  return (
    <SLoginScreenIndex>
       <Card
         shadow="sm" 
         p="lg" 
         radius="md" 
         withBorder
         style={{ display:"flex",justifyContent:"center",alignItems:"center",flexDirection:"column" }}
       >
          <Input 
              placeholder="email" 
              onChange={handleMailIdChange} 
              style={{minWidth : "250px", marginBottom:"15px"}} 
              type={"email"}
           />
           <Button onClick={handleSendAuthOTP}>Send OTP</Button>
       </Card>
    </SLoginScreenIndex>
  )
}

export default LoginScreenIndex