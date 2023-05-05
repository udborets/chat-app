import Head from "next/head";
import { FC } from "react";

import type { IPageLayoutProps } from "./models";

const PageLayout: FC<IPageLayoutProps> = ({ children, title, className }) => {
  return (
    <>
      <Head>
        <title>{title ?? ""}</title>
      </Head>
      <main className={className}>
        {children}
      </main>
    </>
  )
}

export default PageLayout