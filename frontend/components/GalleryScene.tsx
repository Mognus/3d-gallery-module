"use client";

import { Suspense } from "react";
import { Canvas } from "@react-three/fiber";
import { OrbitControls } from "@react-three/drei";
import { PictureFrame } from "./PictureFrame";

export function GalleryScene() {
    return (
        <div className="w-full h-full">
            <Canvas camera={{ position: [0, 0, 20], fov: 40 }}>
                <ambientLight intensity={0.5} />
                <directionalLight position={[5, 5, 5]} intensity={1} />
                <Suspense fallback={null}>
                    <PictureFrame scale={3} imageUrl="https://picsum.photos/400/300" />
                </Suspense>
                <OrbitControls
                    enablePan={false}
                    minDistance={0.5}
                    maxDistance={5}
                />
            </Canvas>
        </div>
    );
}
