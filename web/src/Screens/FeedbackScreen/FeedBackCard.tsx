import { Card, Avatar, Spoiler, Text } from "@mantine/core"
import styled from "styled-components"
import { MFeedBack } from "../../Logic/Modal/MFeedBack"


interface IFeedBackCard{
  feedBack : MFeedBack
}

function FeedBackCard(props:IFeedBackCard) {
  return (
    <Card shadow="lg" p="lg" radius="md" withBorder mb={8} style={{minWidth:"300px",margin:"5px",flexGrow:"1"}}>
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
          {props.feedBack.content}
      </Spoiler>
      <div style={{
        display : "flex",
        justifyContent : "flex-end"
      }}>
          <Text>View Image</Text>
      </div>
  </Card>
  )
}

export default FeedBackCard