"use client";

import { Canvas } from "@react-three/fiber";
import { OrbitControls } from "@react-three/drei";
import { PictureFrame } from "./PictureFrame";

export function GalleryScene() {
    return (
        <div className="w-full h-full">
            <Canvas camera={{ position: [0, 0, 1.5], fov: 50 }}>
                <ambientLight intensity={0.5} />
                <directionalLight position={[5, 5, 5]} intensity={1} />
                <PictureFrame scale={3} />
                <OrbitControls
                    enablePan={false}
                    minDistance={0.5}
                    maxDistance={5}
                />
            </Canvas>
        </div>
    );
}
