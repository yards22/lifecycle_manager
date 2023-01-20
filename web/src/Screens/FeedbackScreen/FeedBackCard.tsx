import { Card, Avatar, Spoiler, Text, Title, Button, Modal, Textarea } from "@mantine/core";
import styled from "styled-components";
import sAgo from "s-ago";
import { MFeedBack } from "../../Logic/Model/MFeedBack";
import FeedBackMedia from "./FeedBackMedia"
import { useState } from "react";
import { useStores } from "../../Logic/Providers/StateProvider";

interface IFeedBackCard {
  feedBack: MFeedBack;
}



function FeedBackCard(props: IFeedBackCard) {
  const stores = useStores();
  const [showModal, setShowModal] = useState(false);
  const [content, setContent] = useState(props.feedBack.comment);

  return (
    <>
       <Modal
        opened={showModal}
        onClose={() => setShowModal(false)}
        title="Bug Fix"
      >
    <Textarea
      placeholder="Say that it's Resolved"
      label="Your comment"
      onChange={(e)=>setContent(e.target.value)}
      withAsterisk
    />
    <div style={{
          paddingTop:10,
          display: "flex",
          justifyContent: "flex-end",
        }}>
      <Button color="green" onClick = {()=>{
        stores.feedBackStore.PostFeedBackComments(!props.feedBack.status,content,props.feedBack.feedback_id).then(()=>{
          setShowModal(false);
        })
        .catch((err)=>{
           throw err
        })
      }}>
        Fixed
      </Button>
    </div>
    </Modal>
    <Card
      shadow="lg"
      p="lg"
      radius="md"
      withBorder
      mb={8}
      style={{ minWidth: "300px", margin: "5px", flexGrow: "1" }}
    >
      <div
        style={{
          display: "flex",
          alignItems: "center",
          justifyContent: "space-between",
        }}
      >
        <div
          style={{
            display: "flex",
            alignItems: "center",
          }}
        >
         <Text fs="italic">@gmail.com</Text>
        </div>
        <Title
            order={6}
            color="dimmed"
            style={{
              marginLeft: "10px",
              fontWeight: "300",
              padding: "0",
              marginTop: "0",
            }}
          >
            {(props.feedBack.created_at.toString().split('T')[0])}
          </Title>
      </div>
      <Spoiler maxHeight={80} showLabel="Show more" hideLabel="Hide">
          {props.feedBack.content}
      </Spoiler>
      {props.feedBack.image_uri && (
          <FeedBackMedia media={props.feedBack.image_uri} />
        )}
      <div
        style={{
          display: "flex",
          justifyContent: "flex-end",
        }}
      >
        {!props.feedBack.status ?<Button fw={500} color="red" onClick={(e: any) => {
            e.preventDefault();
            setShowModal(true);
          }}>Bug</Button> : <Button fw={500} color="green">Fixed</Button> }
      </div>
    </Card>
    </>
  );
}

export default FeedBackCard;
