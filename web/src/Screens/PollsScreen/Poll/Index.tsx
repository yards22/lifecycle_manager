import { Card, Indicator, Title, useMantineTheme } from "@mantine/core";
import { Observer } from "mobx-react-lite";
import { X } from "react-feather";
import styled from "styled-components";
import { useStores } from "../../../Logic/Providers/StateProvider";
import PollButtons from "./PollButtons";
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

function PollCardIndex() {
  const { colors } = useMantineTheme();
  const stores = useStores();

  function handleDeleteAPoll(){
    let confirm = window.confirm("Are you sure to delete this poll");
    if(confirm){

    }
  }

  return (
    <Observer>
      {
        ()=>{
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
                        Who are going to win todays match? #INDVSSA
                      </Title>
                    </div>
                    <div
                      style={{ padding: "30px", display: "flex", flexDirection: "column" }}
                    >
                      <PollButtons
                        buttons={[
                          { title: "IND", votes: 91 },
                          { title: "SA", votes: 1 },
                          { title: "Draw", votes: 9 },
                        ]}
                      />
                    </div>
                  </Card>
              </SEventsIndex>
          )
        }
      }
    </Observer>
  );
}

export default PollCardIndex;
