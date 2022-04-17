import { useEffect, useState } from "react";
import styled from "styled-components";
import { getList } from "./api";
import { Button } from "./components/button";
import { Container } from "./components/container";
import { Paragraph, Title } from "./components/typegraphics";

const Page = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
`;

const Row = styled.div`
  display: flex;
  width: 100%;
  padding: 0 2px;
  margin-bottom: 3px;
  align-items: center;

  *:first-child {
    margin-right: 5px;
  }
`;

function App() {
  const [list, setList] = useState<string[]>([]);

  useEffect(() => {
    loadList();
  }, []);

  const loadList = async () => {
    const list = await getList();
    setList(list);
  };

  const renderList = () => {
    return list.map((item) => {
      return (
        <Row>
          <Button>Wake</Button>
          <Paragraph>{item}</Paragraph>
        </Row>
      );
    });
  };

  return (
    <Page>
      <Container>
        <Title>Wake the Wizzard</Title>

        {renderList()}
      </Container>
    </Page>
  );
}

export default App;
