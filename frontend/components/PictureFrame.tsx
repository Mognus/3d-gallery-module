"use client";

import { useEffect } from "react";
import { useGLTF, useTexture } from "@react-three/drei";
import type { Mesh } from "three";
import { MeshStandardMaterial, SRGBColorSpace } from "three";

const MODEL_PATH = "/models/fancy-picture-frame.glb";
const CANVAS_MESH_NAME = "Plane005";

interface PictureFrameProps {
    imageUrl?: string;
    scale?: number;
}

export function PictureFrame({ imageUrl, scale = 3 }: PictureFrameProps) {
    const { scene } = useGLTF(MODEL_PATH);

    return (
        <group scale={scale}>
            <primitive object={scene} />
            {imageUrl && <CanvasTexture scene={scene} imageUrl={imageUrl} />}
        </group>
    );
}

function CanvasTexture({ scene, imageUrl }: { scene: import("three").Object3D; imageUrl: string }) {
    const texture = useTexture(imageUrl);

    useEffect(() => {
        // GLTF expects flipY=false (UV origin top-left), Three.js defaults to true
        texture.colorSpace = SRGBColorSpace;
        texture.flipY = false;
        // Rotate 180° to match Blender UV orientation
        texture.center.set(0.5, 0.5);
        texture.rotation = Math.PI;

        scene.traverse((child) => {
            if ((child as Mesh).isMesh) {
                const mesh = child as Mesh;
                const mat = mesh.material as MeshStandardMaterial;
                // Find the canvas surface by mesh name or material name
                if (mesh.name === CANVAS_MESH_NAME || mat.name?.includes("canvas")) {
                    const newMat = mat.clone();
                    newMat.map = texture;
                    newMat.color.set(0xffffff);
                    newMat.needsUpdate = true;
                    mesh.material = newMat;
                }
            }
        });
    }, [scene, texture]);

    return null;
}

useGLTF.preload(MODEL_PATH);
