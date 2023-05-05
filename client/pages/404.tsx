import Link from "next/link";
import { FC } from "react";

import PageLayout from "@/layouts/PageLayout/PageLayout";

const NotFoundPage: FC = () => {
  return (
    <PageLayout className="w-full h-full p-2 grid place-content-center">
      <div className="flex flex-col w-fit h-fit gap-4">
        <h1>Page not found</h1>
        <Link href="/">Go home</Link>
      </div>
    </PageLayout>
  )
}

export default NotFoundPage; 