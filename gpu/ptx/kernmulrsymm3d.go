package ptx

//This file is auto-generated. Editing is futile.

func init() { Code["kernmulrsymm3d"] = KERNMULRSYMM3D }

const KERNMULRSYMM3D = `
//
// Generated by NVIDIA NVVM Compiler
// Compiler built on Sat Sep 22 02:35:14 2012 (1348274114)
// Cuda compilation tools, release 5.0, V0.2.1221
//

.version 3.1
.target sm_30
.address_size 64

	.file	1 "/tmp/tmpxft_00007112_00000000-9_kernmulrsymm3d.cpp3.i"
	.file	2 "/home/arne/src/code.google.com/p/nimble-cube/gpu/ptx/kernmulrsymm3d.cu"

.visible .entry kernmulRSymm3D(
	.param .u64 kernmulRSymm3D_param_0,
	.param .u64 kernmulRSymm3D_param_1,
	.param .u64 kernmulRSymm3D_param_2,
	.param .u64 kernmulRSymm3D_param_3,
	.param .u64 kernmulRSymm3D_param_4,
	.param .u64 kernmulRSymm3D_param_5,
	.param .u64 kernmulRSymm3D_param_6,
	.param .u64 kernmulRSymm3D_param_7,
	.param .u64 kernmulRSymm3D_param_8,
	.param .u32 kernmulRSymm3D_param_9,
	.param .u32 kernmulRSymm3D_param_10,
	.param .u32 kernmulRSymm3D_param_11
)
{
	.reg .pred 	%p<7>;
	.reg .s32 	%r<55>;
	.reg .f32 	%f<31>;
	.reg .s64 	%rd<34>;


	ld.param.u64 	%rd10, [kernmulRSymm3D_param_0];
	ld.param.u64 	%rd11, [kernmulRSymm3D_param_1];
	ld.param.u64 	%rd12, [kernmulRSymm3D_param_2];
	ld.param.u64 	%rd9, [kernmulRSymm3D_param_3];
	ld.param.u64 	%rd13, [kernmulRSymm3D_param_4];
	ld.param.u64 	%rd14, [kernmulRSymm3D_param_5];
	ld.param.u64 	%rd15, [kernmulRSymm3D_param_6];
	ld.param.u64 	%rd16, [kernmulRSymm3D_param_7];
	ld.param.u64 	%rd17, [kernmulRSymm3D_param_8];
	ld.param.u32 	%r15, [kernmulRSymm3D_param_9];
	ld.param.u32 	%r16, [kernmulRSymm3D_param_10];
	ld.param.u32 	%r17, [kernmulRSymm3D_param_11];
	cvta.to.global.u64 	%rd1, %rd12;
	cvta.to.global.u64 	%rd2, %rd11;
	cvta.to.global.u64 	%rd3, %rd10;
	cvta.to.global.u64 	%rd4, %rd17;
	cvta.to.global.u64 	%rd5, %rd16;
	cvta.to.global.u64 	%rd6, %rd15;
	cvta.to.global.u64 	%rd7, %rd14;
	cvta.to.global.u64 	%rd8, %rd13;
	.loc 2 36 1
	mov.u32 	%r18, %ntid.y;
	mov.u32 	%r19, %ctaid.y;
	mov.u32 	%r1, %tid.y;
	mad.lo.s32 	%r20, %r18, %r19, %r1;
	.loc 2 37 1
	mov.u32 	%r21, %ntid.x;
	mov.u32 	%r22, %ctaid.x;
	mov.u32 	%r23, %tid.x;
	mad.lo.s32 	%r24, %r21, %r22, %r23;
	.loc 2 39 1
	setp.lt.s32 	%p1, %r24, %r17;
	setp.lt.s32 	%p2, %r20, %r16;
	and.pred  	%p3, %p2, %p1;
	.loc 2 48 1
	setp.gt.s32 	%p4, %r15, 0;
	.loc 2 39 1
	and.pred  	%p5, %p3, %p4;
	@!%p5 bra 	BB0_3;
	bra.uni 	BB0_1;

BB0_1:
	.loc 2 50 1
	mad.lo.s32 	%r51, %r17, %r20, %r24;
	shl.b32 	%r52, %r51, 1;
	add.s32 	%r53, %r52, 1;
	mul.lo.s32 	%r5, %r17, %r16;
	shl.b32 	%r6, %r5, 1;
	mov.u32 	%r54, 0;
	cvta.to.global.u64 	%rd18, %rd9;

BB0_2:
	.loc 2 51 1
	mul.wide.s32 	%rd19, %r51, 4;
	add.s64 	%rd20, %rd18, %rd19;
	.loc 2 52 1
	add.s64 	%rd21, %rd8, %rd19;
	ld.global.f32 	%f1, [%rd21];
	.loc 2 53 1
	add.s64 	%rd22, %rd7, %rd19;
	ld.global.f32 	%f2, [%rd22];
	.loc 2 54 1
	add.s64 	%rd23, %rd6, %rd19;
	ld.global.f32 	%f3, [%rd23];
	.loc 2 55 1
	add.s64 	%rd24, %rd5, %rd19;
	.loc 2 56 1
	add.s64 	%rd25, %rd4, %rd19;
	.loc 2 59 1
	mul.wide.s32 	%rd26, %r52, 4;
	add.s64 	%rd27, %rd3, %rd26;
	.loc 2 60 1
	mul.wide.s32 	%rd28, %r53, 4;
	add.s64 	%rd29, %rd3, %rd28;
	ld.global.f32 	%f4, [%rd29];
	.loc 2 61 1
	add.s64 	%rd30, %rd2, %rd26;
	.loc 2 62 1
	add.s64 	%rd31, %rd2, %rd28;
	ld.global.f32 	%f5, [%rd31];
	.loc 2 63 1
	add.s64 	%rd32, %rd1, %rd26;
	.loc 2 64 1
	add.s64 	%rd33, %rd1, %rd28;
	ld.global.f32 	%f6, [%rd33];
	.loc 2 59 1
	ld.global.f32 	%f7, [%rd27];
	.loc 2 51 1
	ld.global.f32 	%f8, [%rd20];
	.loc 2 61 1
	ld.global.f32 	%f9, [%rd30];
	.loc 2 56 1
	ld.global.f32 	%f10, [%rd25];
	.loc 2 66 1
	mul.f32 	%f11, %f9, %f10;
	fma.rn.f32 	%f12, %f7, %f8, %f11;
	.loc 2 63 1
	ld.global.f32 	%f13, [%rd32];
	.loc 2 55 1
	ld.global.f32 	%f14, [%rd24];
	.loc 2 66 1
	fma.rn.f32 	%f15, %f13, %f14, %f12;
	st.global.f32 	[%rd27], %f15;
	.loc 2 67 1
	mul.f32 	%f16, %f5, %f10;
	fma.rn.f32 	%f17, %f4, %f8, %f16;
	fma.rn.f32 	%f18, %f6, %f14, %f17;
	st.global.f32 	[%rd29], %f18;
	.loc 2 68 1
	mul.f32 	%f19, %f9, %f1;
	fma.rn.f32 	%f20, %f7, %f10, %f19;
	fma.rn.f32 	%f21, %f13, %f3, %f20;
	st.global.f32 	[%rd30], %f21;
	.loc 2 69 1
	mul.f32 	%f22, %f5, %f1;
	fma.rn.f32 	%f23, %f4, %f10, %f22;
	fma.rn.f32 	%f24, %f6, %f3, %f23;
	st.global.f32 	[%rd31], %f24;
	.loc 2 70 1
	mul.f32 	%f25, %f9, %f3;
	fma.rn.f32 	%f26, %f7, %f14, %f25;
	fma.rn.f32 	%f27, %f13, %f2, %f26;
	st.global.f32 	[%rd32], %f27;
	.loc 2 71 1
	mul.f32 	%f28, %f5, %f3;
	fma.rn.f32 	%f29, %f4, %f14, %f28;
	fma.rn.f32 	%f30, %f6, %f2, %f29;
	st.global.f32 	[%rd33], %f30;
	.loc 2 48 1
	add.s32 	%r53, %r53, %r6;
	add.s32 	%r52, %r52, %r6;
	add.s32 	%r51, %r51, %r5;
	.loc 2 48 18
	add.s32 	%r54, %r54, 1;
	.loc 2 48 1
	setp.lt.s32 	%p6, %r54, %r15;
	@%p6 bra 	BB0_2;

BB0_3:
	.loc 2 73 2
	ret;
}


`
