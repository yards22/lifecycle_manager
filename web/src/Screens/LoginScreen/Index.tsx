import { Button, Card, Input, Loader } from '@mantine/core'
import React, { useEffect, useState } from 'react'
import styled from 'styled-components'
import { useStores } from '../../Logic/Providers/StateProvider'
import { Navigate, useNavigate } from 'react-router-dom'

const SLoginScreenIndex = styled.div`
    display: flex;
    width: 100%;
    height: 100%;
    justify-content: center;
    align-items: center;

`

const AUTH_INITIAL = 0;
const CHECKING_AUTH = 1;
const CHECKED_AUTH_LOGGED_IN = 2;

function LoginScreenIndex() {
  const store = useStores();
  const [mailId,setMailId] = useState("")
  const navigate = useNavigate()
  const [otp,setOTP] = useState("")
  const [authStage,setAuthStage] = useState(CHECKING_AUTH);
  const [currentState , setCurrentState] = useState(1)


useEffect(()=>{
  if (
    !store.authStore.userMailId &&
    (store.authStore.token == null || store.authStore.token === "")
  ){
    setAuthStage(AUTH_INITIAL);
  }
  else if (store.authStore.token) {
    // check if logged in using token
    setAuthStage(CHECKING_AUTH);
    store.authStore
      .CheckIfLogin()
      .then(() => {
        setAuthStage(CHECKED_AUTH_LOGGED_IN);
      })
      .catch((err) => {
        setAuthStage(AUTH_INITIAL);
      });
  }
},[])

useEffect(() => {
  if (store.authStore.token) {
    store.authStore
      .CheckIfLogin()
      .then(() => {
        setAuthStage(CHECKED_AUTH_LOGGED_IN);
      })
      .catch((err) => {});
  }
}, []);

  function handleMailIdChange(e:any){
    setMailId(e.target.value)
  }

  function handleOTPChange(e:any){
    setOTP(e.target.value)
  }

  async function handleSendAuthOTP(){
    const res = await store.authStore.SendAuthOTP(mailId)
    if(res === 200){

      setCurrentState(2)
    }
  }

  async function handleVerifyOTP() {
    const res = await store.authStore.VerifyAuthOTP(mailId,otp)
    console.log("mail_ID",res)
    if(res===200){
      store.authStore.SetIsUserLoggedIn(true)
      navigate({
        pathname : "/feedback"
      })
       
    }
  }

  if (authStage === CHECKING_AUTH) {
    return (
      <section
        style={{
          height: "100vh",
          width: "100vh",
          display: "flex",
          justifyContent: "center",
          alignItems: "center"
        }}
      >
        <Loader variant="bars" size={"sm"} />
      </section>
    );
  } else if (authStage === CHECKED_AUTH_LOGGED_IN)
    return <Navigate to="/feedback" />;

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