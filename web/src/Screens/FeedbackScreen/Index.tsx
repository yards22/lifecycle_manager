import { Avatar, Card, Spoiler, Text } from '@mantine/core'
import React from 'react'
import styled from 'styled-components'
import { DummyTestDataFeedBack } from '../../Data/DummyFeedBackData'

const SFeedBackScreenIndex = styled.div`
   display: flex;
   width: 100%;
   display: flex;
   flex-direction: column;
`

function FeedBackScreenIndex() {
  return (
    <SFeedBackScreenIndex>
      {
        DummyTestDataFeedBack.map((each)=>{
          return(
             <Card key={each.key} shadow="lg" p="lg" radius="md" withBorder mb={8}>
                <div style={{
                  display : "flex",
                  alignItems : "center",
                  justifyContent : "space-between"
                }}>
                   <div style={{
                    display : "flex",
                    alignItems : "center"
                  }}>
                      <Avatar m={6} ml={0} size={50}/>
                      <Text>@username</Text>
                   </div>
                   <div>
                      Date : 21/10/2022
                   </div>
                </div>
                <Spoiler maxHeight={80} showLabel="Show more" hideLabel="Hide">
                    {each.content}
                </Spoiler>
                <div style={{
                  display : "flex",
                  justifyContent : "flex-end"
                }}>
                    <Text>View Image</Text>
                </div>
             </Card>
          )
        })
      }
    </SFeedBackScreenIndex>
  )
}

export default FeedBackScreenIndex