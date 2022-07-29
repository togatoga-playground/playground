import { GetStaticPaths, GetStaticProps, NextPage } from "next";
import Head from "next/head";
import { useRouter } from "next/router";
type PostPros = {
  id: string;
};

const Post: NextPage<PostPros> = (props) => {
  const { id } = props;
  const router = useRouter();
  if (router.isFallback) {
    return <div>Loading...</div>
  }
  return (
    <div>
      <Head>
        <title>Create Next App</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main>
        <p>This site is generated</p>
        <p>{`This page corresponds to /posts/${id}`}</p>
      </main>
    </div>
  );
};

export const getStaticPaths: GetStaticPaths = async () => {
  const paths = [
    {
      params: {
        id: "1",
      },
    },
    {
      params: {
        id: "2",
      },
    },
    {
      params: {
        id: "3",
      },
    },
    {
      params: {
        id: "4",
      },
    },
  ];
  return { paths, fallback: false };
};

export const getStaticProps: GetStaticProps<PostPros> = async (context) => {
  const params = context.params!['id']!;
  
  console.log(params);
  const id = Array.isArray(params)
    ? params[0]
    : params;
  return {
    props: {
      id,
    },
  };
};

export default Post;
