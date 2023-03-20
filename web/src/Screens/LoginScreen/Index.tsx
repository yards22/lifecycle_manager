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
  const [otp,setOTP] = useState("")
  const [currentState , setCurrentState] = useState(1)

  function handleMailIdChange(e:any){
    setMailId(e.target.value)
  }

  function handleOTPChange(e:any){
    setOTP(e.target.value)
  }

  async function handleSendAuthOTP(){
    const res = await stores.authStore.SendAuthOTP(mailId)
    if(res === 200){
      setCurrentState(2)
    }
  }

  async function handleVerifyOTP() {
    const res = await stores.authStore.VerifyAuthOTP(mailId,otp)
    console.log("mail_ID",res)
    if(res===200){
      stores.authStore.SetIsUserLoggedIn(true)
      window.location.reload()
    }
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
         { currentState === 1 &&
          <>         
             <Input 
                placeholder="email" 
                onChange={handleMailIdChange} 
                style={{minWidth : "250px", marginBottom:"15px"}} 
              />
              <Button onClick={handleSendAuthOTP}>Send OTP</Button>
           </>
          }
          {
            currentState === 2 &&
            <>
              <Input 
                placeholder="OTP" 
                onChange={handleOTPChange} 
                style={{minWidth : "250px", marginBottom:"15px"}} 
              />
              <Button onClick={handleVerifyOTP}>Verify OTP</Button>
            </>
          }
       </Card>
    </SLoginScreenIndex>
  )
}

export default LoginScreenIndex