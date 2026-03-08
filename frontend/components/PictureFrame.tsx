"use client";

import { useEffect } from "react";
import { useGLTF, useTexture } from "@react-three/drei";
import type { Mesh } from "three";
import { MeshStandardMaterial, SRGBColorSpace } from "three";

interface PictureFrameProps {
    modelUrl: string;
    canvasMeshName?: string;
    imageUrl?: string;
    scale?: number;
}

export function PictureFrame({ modelUrl, canvasMeshName, imageUrl, scale = 3 }: PictureFrameProps) {
    const { scene } = useGLTF(modelUrl);

    return (
        <group scale={scale}>
            <primitive object={scene} />
            {imageUrl && canvasMeshName && (
                <CanvasTexture scene={scene} imageUrl={imageUrl} canvasMeshName={canvasMeshName} />
            )}
        </group>
    );
}

function CanvasTexture({
    scene,
    imageUrl,
    canvasMeshName,
}: {
    scene: import("three").Object3D;
    imageUrl: string;
    canvasMeshName: string;
}) {
    const texture = useTexture(imageUrl);

    useEffect(() => {
        texture.colorSpace = SRGBColorSpace;
        texture.flipY = false;
        texture.center.set(0.5, 0.5);
        texture.rotation = Math.PI;

        scene.traverse((child) => {
            if ((child as Mesh).isMesh) {
                const mesh = child as Mesh;
                const mat = mesh.material as MeshStandardMaterial;
                if (mesh.name === canvasMeshName || mat.name?.includes("canvas")) {
                    const newMat = mat.clone();
                    newMat.map = texture;
                    newMat.color.set(0xffffff);
                    newMat.needsUpdate = true;
                    mesh.material = newMat;
                }
            }
        });
    }, [scene, texture, canvasMeshName]);

    return null;
}
