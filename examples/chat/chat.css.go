package main

var chat_css = []byte{
	0x23, 0x6d, 0x73, 0x67, 0x20, 0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x62,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x3a, 0x20, 0x31, 0x70, 0x78, 0x20, 0x73,
	0x6f, 0x6c, 0x69, 0x64, 0x20, 0x67, 0x72, 0x61, 0x79, 0x3b, 0x0a, 0x20,
	0x20, 0x20, 0x20, 0x70, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x20,
	0x30, 0x2e, 0x35, 0x65, 0x6d, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x6d,
	0x61, 0x78, 0x2d, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x20, 0x38,
	0x30, 0x25, 0x3b, 0x0a, 0x7d, 0x0a, 0x0a, 0x23, 0x6d, 0x73, 0x67, 0x20,
	0x70, 0x20, 0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x6d, 0x61, 0x72, 0x67,
	0x69, 0x6e, 0x3a, 0x20, 0x30, 0x2e, 0x31, 0x65, 0x6d, 0x3b, 0x0a, 0x7d,
	0x0a, 0x0a, 0x23, 0x6d, 0x73, 0x67, 0x20, 0x73, 0x70, 0x61, 0x6e, 0x20,
	0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e,
	0x2d, 0x72, 0x69, 0x67, 0x68, 0x74, 0x3a, 0x20, 0x30, 0x2e, 0x35, 0x65,
	0x6d, 0x3b, 0x0a, 0x7d, 0x0a, 0x0a, 0x23, 0x6d, 0x73, 0x67, 0x20, 0x2e,
	0x74, 0x69, 0x6d, 0x65, 0x20, 0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x66,
	0x6f, 0x6e, 0x74, 0x2d, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x3a, 0x20,
	0x6d, 0x6f, 0x6e, 0x6f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x3b, 0x0a, 0x20,
	0x20, 0x20, 0x20, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x73, 0x69, 0x7a, 0x65,
	0x3a, 0x20, 0x38, 0x30, 0x25, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x63,
	0x6f, 0x6c, 0x6f, 0x72, 0x3a, 0x20, 0x23, 0x34, 0x30, 0x34, 0x30, 0x34,
	0x30, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x66, 0x6c, 0x6f, 0x61, 0x74,
	0x3a, 0x20, 0x72, 0x69, 0x67, 0x68, 0x74, 0x3b, 0x0a, 0x7d, 0x0a, 0x0a,
	0x23, 0x6d, 0x73, 0x67, 0x20, 0x2e, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x7b,
	0x0a, 0x20, 0x20, 0x20, 0x20, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x3a, 0x20,
	0x62, 0x6c, 0x75, 0x65, 0x3b, 0x0a, 0x7d, 0x0a, 0x0a, 0x23, 0x6d, 0x73,
	0x67, 0x20, 0x2e, 0x66, 0x72, 0x6f, 0x6d, 0x3a, 0x61, 0x66, 0x74, 0x65,
	0x72, 0x20, 0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x3a, 0x20, 0x22, 0x3a, 0x22, 0x3b, 0x0a, 0x7d, 0x0a,
	0x0a, 0x23, 0x6d, 0x73, 0x67, 0x20, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x20, 0x7b, 0x0a, 0x7d, 0x0a, 0x0a, 0x66, 0x6f, 0x72, 0x6d,
	0x5b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x22, 0x6d, 0x73, 0x67, 0x66, 0x6f,
	0x72, 0x6d, 0x22, 0x5d, 0x20, 0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x6d,
	0x61, 0x78, 0x2d, 0x77, 0x69, 0x64, 0x74, 0x68, 0x3a, 0x20, 0x39, 0x35,
	0x25, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x3a, 0x20, 0x2d, 0x77, 0x65, 0x62, 0x6b, 0x69, 0x74, 0x2d,
	0x66, 0x6c, 0x65, 0x78, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x3a, 0x20, 0x2d, 0x6d, 0x6f, 0x7a, 0x2d,
	0x66, 0x6c, 0x65, 0x78, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x3a, 0x20, 0x2d, 0x6d, 0x73, 0x2d, 0x66,
	0x6c, 0x65, 0x78, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x3a, 0x20, 0x66, 0x6c, 0x65, 0x78, 0x3b, 0x0a,
	0x20, 0x20, 0x20, 0x20, 0x2d, 0x77, 0x65, 0x62, 0x6b, 0x69, 0x74, 0x2d,
	0x66, 0x6c, 0x65, 0x78, 0x2d, 0x66, 0x6c, 0x6f, 0x77, 0x3a, 0x20, 0x72,
	0x6f, 0x77, 0x20, 0x6e, 0x6f, 0x77, 0x72, 0x61, 0x70, 0x3b, 0x0a, 0x20,
	0x20, 0x20, 0x20, 0x2d, 0x6d, 0x6f, 0x7a, 0x2d, 0x66, 0x6c, 0x65, 0x78,
	0x2d, 0x66, 0x6c, 0x6f, 0x77, 0x3a, 0x20, 0x72, 0x6f, 0x77, 0x20, 0x6e,
	0x6f, 0x77, 0x72, 0x61, 0x70, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x2d,
	0x6d, 0x73, 0x2d, 0x66, 0x6c, 0x65, 0x78, 0x2d, 0x66, 0x6c, 0x6f, 0x77,
	0x3a, 0x20, 0x72, 0x6f, 0x77, 0x20, 0x6e, 0x6f, 0x77, 0x72, 0x61, 0x70,
	0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x2d, 0x6d, 0x73, 0x2d, 0x66, 0x6c,
	0x65, 0x78, 0x2d, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x3a, 0x20, 0x72, 0x6f, 0x77, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x2d,
	0x6d, 0x73, 0x2d, 0x66, 0x6c, 0x65, 0x78, 0x2d, 0x77, 0x72, 0x61, 0x70,
	0x3a, 0x20, 0x6e, 0x6f, 0x77, 0x72, 0x61, 0x70, 0x3b, 0x0a, 0x20, 0x20,
	0x20, 0x20, 0x66, 0x6c, 0x65, 0x78, 0x2d, 0x66, 0x6c, 0x6f, 0x77, 0x3a,
	0x20, 0x72, 0x6f, 0x77, 0x20, 0x6e, 0x6f, 0x77, 0x72, 0x61, 0x70, 0x3b,
	0x0a, 0x7d, 0x0a, 0x0a, 0x66, 0x6f, 0x72, 0x6d, 0x5b, 0x6e, 0x61, 0x6d,
	0x65, 0x3d, 0x22, 0x6d, 0x73, 0x67, 0x66, 0x6f, 0x72, 0x6d, 0x22, 0x5d,
	0x20, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x20, 0x7b, 0x0a, 0x20, 0x20, 0x20,
	0x20, 0x2d, 0x77, 0x65, 0x62, 0x6b, 0x69, 0x74, 0x2d, 0x66, 0x6c, 0x65,
	0x78, 0x3a, 0x20, 0x31, 0x20, 0x61, 0x75, 0x74, 0x6f, 0x3b, 0x0a, 0x20,
	0x20, 0x20, 0x20, 0x2d, 0x6d, 0x6f, 0x7a, 0x2d, 0x66, 0x6c, 0x65, 0x78,
	0x3a, 0x20, 0x31, 0x20, 0x61, 0x75, 0x74, 0x6f, 0x3b, 0x0a, 0x20, 0x20,
	0x20, 0x20, 0x2d, 0x6d, 0x73, 0x2d, 0x66, 0x6c, 0x65, 0x78, 0x3a, 0x20,
	0x31, 0x20, 0x61, 0x75, 0x74, 0x6f, 0x3b, 0x0a, 0x20, 0x20, 0x20, 0x20,
	0x66, 0x6c, 0x65, 0x78, 0x3a, 0x20, 0x31, 0x20, 0x61, 0x75, 0x74, 0x6f,
	0x3b, 0x0a, 0x7d, 0x0a, 0x0a, 0x66, 0x6f, 0x72, 0x6d, 0x5b, 0x6e, 0x61,
	0x6d, 0x65, 0x3d, 0x22, 0x6d, 0x73, 0x67, 0x66, 0x6f, 0x72, 0x6d, 0x22,
	0x5d, 0x20, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5b, 0x6e, 0x61, 0x6d, 0x65,
	0x3d, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x5d, 0x20, 0x7b, 0x0a, 0x20,
	0x20, 0x20, 0x20, 0x6d, 0x61, 0x78, 0x2d, 0x77, 0x69, 0x64, 0x74, 0x68,
	0x3a, 0x20, 0x38, 0x65, 0x6d, 0x3b, 0x0a, 0x7d, 0x0a, 0x0a, 0x66, 0x6f,
	0x72, 0x6d, 0x5b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x22, 0x6d, 0x73, 0x67,
	0x66, 0x6f, 0x72, 0x6d, 0x22, 0x5d, 0x20, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x5b, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x22, 0x73, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x22, 0x5d, 0x20, 0x7b, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x2d, 0x77,
	0x65, 0x62, 0x6b, 0x69, 0x74, 0x2d, 0x66, 0x6c, 0x65, 0x78, 0x3a, 0x20,
	0x30, 0x20, 0x30, 0x20, 0x61, 0x75, 0x74, 0x6f, 0x3b, 0x0a, 0x20, 0x20,
	0x20, 0x20, 0x2d, 0x6d, 0x6f, 0x7a, 0x2d, 0x66, 0x6c, 0x65, 0x78, 0x3a,
	0x20, 0x30, 0x20, 0x30, 0x20, 0x61, 0x75, 0x74, 0x6f, 0x3b, 0x0a, 0x20,
	0x20, 0x20, 0x20, 0x2d, 0x6d, 0x73, 0x2d, 0x66, 0x6c, 0x65, 0x78, 0x3a,
	0x20, 0x30, 0x20, 0x30, 0x20, 0x61, 0x75, 0x74, 0x6f, 0x3b, 0x0a, 0x20,
	0x20, 0x20, 0x20, 0x66, 0x6c, 0x65, 0x78, 0x3a, 0x20, 0x30, 0x20, 0x30,
	0x20, 0x61, 0x75, 0x74, 0x6f, 0x3b, 0x0a, 0x7d, 0x0a,
}
