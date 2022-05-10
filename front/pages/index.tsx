import type { NextPage } from 'next'
import dynamic from 'next/dynamic'
import React from 'react'
import Head from 'next/head'
import Image from 'next/image'
import styles from '../styles/Home.module.css'

import Test from '../components/Test'
import 'leaflet/dist/leaflet.css';

const Home: NextPage = () => {
  const Map = React.useMemo(
    () => dynamic(() => import("../components/Test"), {
        loading: () => <p>loading...</p>,
        ssr: false,
      }), []
  );
  return <Map />;
}

export default Home
