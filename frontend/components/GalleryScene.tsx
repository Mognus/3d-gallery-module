"use client";

import { Suspense } from "react";
import { Canvas } from "@react-three/fiber";
import { OrbitControls } from "@react-three/drei";
import { PictureFrame } from "./PictureFrame";

interface GallerySceneProps {
    modelUrl: string;
    canvasMeshName?: string;
    imageUrl?: string;
    scale?: number;
    className?: string;
}

export function GalleryScene({ modelUrl, canvasMeshName, imageUrl, scale = 3, className }: GallerySceneProps) {
    return (
        <div className={className ?? "w-full h-full"}>
            <Canvas camera={{ position: [0, 0, 20], fov: 40 }}>
                <ambientLight intensity={0.5} />
                <directionalLight position={[5, 5, 5]} intensity={1} />
                <Suspense fallback={null}>
                    <PictureFrame
                        modelUrl={modelUrl}
                        canvasMeshName={canvasMeshName}
                        imageUrl={imageUrl}
                        scale={scale}
                    />
                </Suspense>
                <OrbitControls enablePan={false} minDistance={0.5} maxDistance={5} />
            </Canvas>
        </div>
    );
}
