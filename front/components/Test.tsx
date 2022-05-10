import React, { useState } from "react";
import { MapContainer, TileLayer, Marker, Popup, GeoJSON } from "react-leaflet";
import Leaflet from "leaflet";
import geodata from "../public/prefectures.json";

import { openReverseGeocoder as geocoder } from "@geolonia/open-reverse-geocoder";

Leaflet.Icon.Default.imagePath =
  "//cdnjs.cloudflare.com/ajax/libs/leaflet/1.3.1/images/";

const onEachFeature = (feature, layer) => {
  layer.on({
    click: async (e) => {
      const { lat, lng } = e.latlng;
      const { prefecture, city } = await geocoder([lng, lat]);
    },
  });
};


const Test = () => {
  const position = [35.712323, 139.9469441];
  const fetchNews = (q) => {
    fetch('')
    }
  return (
    <MapContainer center={position} zoom={13} scrollWheelZoom={false}>
      <TileLayer
        attribution={
          '&copy; <a href="https://maps.gsi.go.jp/development/ichiran.html">国土地理院</a>'
        }
        url="https://cyberjapandata.gsi.go.jp/xyz/std/{z}/{x}/{y}.png"
      />
      <GeoJSON data={geodata} onEachFeature={onEachFeature} />
    </MapContainer>
  );
};
export default Test;
