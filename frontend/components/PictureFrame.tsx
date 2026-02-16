"use client";

import { useMemo } from "react";
import { useGLTF } from "@react-three/drei";

interface PictureFrameProps {
    scale?: number;
}

export function PictureFrame({ scale = 3 }: PictureFrameProps) {
    const { scene } = useGLTF("/models/picture-frame.glb");
    const clone = useMemo(() => scene.clone(), [scene]);

    return <primitive object={clone} scale={scale} />;
}

useGLTF.preload("/models/picture-frame.glb");
