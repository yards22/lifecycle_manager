import { Card, Indicator, Title, useMantineTheme } from "@mantine/core";
import { Observer } from "mobx-react-lite";
import { X } from "react-feather";
import styled from "styled-components";
import { useStores } from "../../Logic/Providers/StateProvider";
import PollButtons from "./PollButtons";
import { useEffect } from "react";
import { MPollButton, MPolls } from "../../Logic/Model/MPolls";
import { string } from "yargs";
const SEventsIndex = styled.div`
  height: 100%;
  position: relative;
  overflow: hidden;
  margin: 5px;
  flex-grow: 1;
`;


const Flag = styled.div`
  position: absolute;
  left: -40px;
  top: 20px;
  z-index: 10;
  width: 150px;
  background: #ff5b5b;
  display: flex;
  justify-content: center;
  color: white;
  font-weight: bold;
  transform: rotate(-45deg);
`;

interface IPolls {
  polls: MPolls;
}

export default function PollCards(props:IPolls) {
  const { colors } = useMantineTheme();
  const stores = useStores();
  useEffect(()=>{
    stores.pollsStore.GetPolls()
  },[]);
  const options= (props.polls.options).split(',')
  
  return (
    <Observer>
      {
        ()=>{
          if (options.length>0){
            return(
                 <SEventsIndex>
                    <Card shadow="lg" p="lg" withBorder style={{ padding: "0" }}>
                      <Flag>Poll</Flag>
                      <div
                        style={{
                          background: colors.blue[8],
                          padding: "10px 10px",
                          minHeight: "50px",
                        }}
                      >
                        <Title order={5} color="white" style={{ marginLeft: "80px" }}>
                          {props.polls.poll_question}
                        </Title>
                      </div>
                      <div
                        style={{ padding: "30px", display: "flex", flexDirection: "column" }}
                      >
                      {
                        options.map((item,index)=>{
                          let x :MPollButton = {title:item,votes:0};
                          return (
                             <p>Pending</p>
                          )
                        }
                        )
                      }
                      </div>
                    </Card>
                </SEventsIndex>
            )
          }
          return (
            <p>Invalid Poll</p>
          )
        }
      }
    </Observer>
  );
}
