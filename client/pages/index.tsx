import { FC } from "react";

import PageLayout from "@/layouts/PageLayout/PageLayout";
import ChatBar from "@/features/ChatBar/ChatBar";
import Chat from "@/features/Chat/Chat";

const HomePage: FC = () => {
  return (
    <PageLayout title="YouChat" className="w-full h-full flex">
      <ChatBar />
      <Chat />
    </PageLayout>
  )
}

export default HomePage;