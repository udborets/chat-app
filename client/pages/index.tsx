import { FC } from "react";

import PageLayout from "@/layouts/PageLayout/PageLayout";
import ChatBar from "@/features/ChatBar/ChatBar";

const HomePage: FC = () => {
  return (
    <PageLayout title="YouChat" className="w-full h-full">
      <ChatBar />
    </PageLayout>
  )
}

export default HomePage;