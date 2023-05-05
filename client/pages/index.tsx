import { FC } from "react";

import PageLayout from "@/layouts/PageLayout/PageLayout";
import ChatBar from "@/features/ChatBar/ChatBar";
import Chat from "@/features/Chat/Chat";
import { ChatBarProps } from "@/features/ChatBar/models";

const testChats: ChatBarProps = {
  chats: [
    {
      companionAvatar: '',
      companionName: "test 1",
      id: Math.random(),
      lastMessage: "Hello!",
    },
    {
      companionAvatar: '',
      companionName: "test 1",
      id: Math.random(),
      lastMessage: "Hello!",
    },
    {
      companionAvatar: '',
      companionName: "test 1",
      id: Math.random(),
      lastMessage: "Hello!",
    },
    {
      companionAvatar: '',
      companionName: "test 1",
      id: Math.random(),
      lastMessage: "Hello!",
    },
    {
      companionAvatar: '',
      companionName: "test 1",
      id: Math.random(),
      lastMessage: "Hello!",
    },
  ]
}

const HomePage: FC = () => {
  return (
    <PageLayout title="YouChat" className="w-full h-full flex">
      <ChatBar {...testChats} />
      <Chat />
    </PageLayout>
  )
}

export default HomePage;