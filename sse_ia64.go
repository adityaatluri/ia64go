
package sse_ia64

/*
 #include<xmmintrin.h>
 void __MM_TRANSPOSE4_PS(__m128 a , __m128 b, __m128 c, __m128 d) {
	_MM_TRANSPOSE4_PS(a, b, c, d);
 }
*/
import "C"
import "unsafe"

func MM_Addps(a C.__m128, b C.__m128) C.__m128 {
	return C._mm_add_ps(a, b)
}

func MM_Addss(a C.__m128, b C.__m128) C.__m128 {
	return C._mm_add_ss(a, b)
}

func MM_Andps(a C.__m128, b C.__m128) C.__m128 {
	return C._mm_and_ps(a, b)
}

func MM_Andnotps(a C.__m128, b C.__m128) C.__m128 {
	return C._mm_andnot_ps(a, b)
}

func MM_Avgpu16(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_avg_pu16(a, b)
}

func MM_Avgpu8(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_avg_pu8(a, b)
}

func MM_Cmpeqps(a C.__m128, b C.__m128) C.__m128 {
	return C._mm_cmpeq_ps(a, b)
}

func MM_Cmpeqss(a C.__m128, b C.__m128) C.__m128 {
	return C._mm_cmpeq_ss(a, b)
}

func MM_Cmpgeps(a C.__m128, b C.__m128) C.__m128 {
	return C._mm_cmpge_ps(a, b)
}

func MM_Cmpgess(a C.__m128, b C.__m128) C.__m128 {
	return C._mm_cmpge_ss(a, b)
}

func MM_Cmpgtps(a C.__m128, b C.__m128) C.__m128 {
	return C._mm_cmpgt_ps(a, b)
}

func MM_Cmpgtss(a C.__m128, b C.__m128) C.__m128 {
	return C._mm_cmpgt_ss(a, b)

}

func MM_Cmpleps(a C.__m128, b C.__m128) C.__m128 {
	return C._mm_cmple_ps(a, b)
}

func MM_Cmpless(a C.__m128, b C.__m128) C.__m128 {
	return C._mm_cmple_ss(a, b)
}

func MM_Cmpltps(a C.__m128, b C.__m128) C.__m128 {
	return C._mm_cmplt_ps(a, b)
}

func MM_Cmpltss(a C.__m128, b C.__m128) C.__m128 {
	return C._mm_cmplt_ss(a, b)
}

func MM_Cmpneqps(a C.__m128, b C.__m128) C.__m128 {
	return C._mm_cmpneq_ps(a, b)
}

func MM_Cmpneqss(a C.__m128, b C.__m128) C.__m128 {
	return C._mm_cmpneq_ss(a, b)
}

func MM_Cmpngeps(a C.__m128, b C.__m128) C.__m128 {
	return C._mm_cmpnge_ps(a, b)
}

func MM_Cmpngess(a C.__m128, b C.__m128) C.__m128 {
	return C._mm_cmpnge_ss(a, b)
}

func MM_Cmpngtps(a C.__m128, b C.__m128) C.__m128 {
	return C._mm_cmpngt_ps(a, b)
}

func MM_Cmpngtss(a C.__m128, b C.__m128) C.__m128 {
	return C._mm_cmpngt_ss(a, b)
}

func MM_Cmpnleps(a C.__m128, b C.__m128) C.__m128 {
	return C._mm_cmpnle_ps(a, b)
}

func MM_Cmpnless(a C.__m128, b C.__m128) C.__m128 {
	return C._mm_cmpnle_ss(a, b)
}

func MM_Cmpnltps(a C.__m128, b C.__m128) C.__m128 {
	return C._mm_cmpnlt_ps(a, b)
}

func MM_Cmpnltss(a C.__m128, b C.__m128) C.__m128 {
	return C._mm_cmpnlt_ss(a, b)
}

func MM_Cmpordps(a C.__m128, b C.__m128) C.__m128 {
	return C._mm_cmpord_ps(a, b)
}

func MM_Cmpordss(a C.__m128, b C.__m128) C.__m128 {
	return C._mm_cmpord_ss(a, b)
}

func MM_Cmupnordps(a C.__m128, b C.__m128) C.__m128 {
	return C._mm_cmpunord_ps(a, b)
}

func MM_Cmpunordss(a C.__m128, b C.__m128) C.__m128 {
	return C._mm_cmpunord_ss(a, b)
}

func MM_Comieqss(a C.__m128, b C.__m128) int {
	return int(C._mm_comieq_ss(a, b))
}

func MM_Comigess(a C.__m128, b C.__m128) int {
	return int(C._mm_comige_ss(a, b))
}

func MM_Comigtss(a C.__m128, b C.__m128) int {
	return int(C._mm_comigt_ss(a, b))
}

func MM_Comiless(a C.__m128, b C.__m128) int {
	return int(C._mm_comile_ss(a, b))
}

func MM_Comiltss(a C.__m128, b C.__m128) int {
	return int(C._mm_comilt_ss(a, b))
}

func MM_Comineqss(a C.__m128, b C.__m128) int {
	return int(C._mm_comineq_ss(a, b))
}

func MM_Cvtpi2ps(a C.__m128, b C.__m64) C.__m128 {
	return C._mm_cvt_pi2ps(a, b)
}

func MM_Cvtps2pi(a C.__m128) C.__m64 {
	return C._mm_cvt_ps2pi(a)
}

func MM_Cvtsi2ss(a C.__m128, b int) C.__m128 {
	return C._mm_cvt_si2ss(a, (C.int)(b))
}

func MM_Cvtss2si(a C.__m128) int {
	return int(C._mm_cvt_ss2si(a))
}

func MM_Cvtpi16ps(a C.__m64) C.__m128 {
	return C._mm_cvtpi16_ps(a)
}

func MM_Cvtpi32ps(a C.__m128, b C.__m64) C.__m128 {
	return C._mm_cvtpi32_ps(a, b)
}

func MM_Cvtpi32x2ps(a C.__m64, b C.__m64) C.__m128 {
	return C._mm_cvtpi32x2_ps(a, b)
}

func MM_Cvtpi8ps(a C.__m64) C.__m128 {
	return C._mm_cvtpi8_ps(a)
}

func MM_Cvtpspi16(a C.__m128) C.__m64 {
	return C._mm_cvtps_pi16(a)
}

func MM_Cvtpspi32(a C.__m128) C.__m64 {
	return C._mm_cvtps_pi32(a)
}

func MM_Cvtpspi8(a C.__m128) C.__m64 {
	return C._mm_cvtps_pi8(a)
}

func MM_Cvtpu16ps(a C.__m64) C.__m128 {
	return C._mm_cvtpu16_ps(a)
}

func MM_Cvtpu8ps(a C.__m64) C.__m128 {
	return C._mm_cvtpu8_ps(a)
}

func MM_Cvtsi32ss(a C.__m128, b C.int) C.__m128 {
	return C._mm_cvtsi32_ss(a, b)
}

func MM_Cvtsi64ss(a C.__m128, b int64) C.__m128 {
	return C._mm_cvtsi64_ss(a, (C.longlong)(b))
}

func MM_Cvtssf32(a C.__m128) float32 {
	return (float32)(C._mm_cvtss_f32(a))
}

func MM_Cvtsssi32(a C.__m128) int {
	return (int)(C._mm_cvtss_si32(a))
}

func MM_Cvtsssi64(a C.__m128) int64 {
	return (int64)(C._mm_cvtss_si64(a))
}

func MM_Cvttps2pi(a C.__m128) C.__m64 {
	return C._mm_cvtt_ps2pi(a)
}

func MM_Cvttss2si(a C.__m128) int {
	return (int)(C._mm_cvtt_ss2si(a))
}

func MM_Cvttpspi32(a C.__m128) C.__m64 {
	return C._mm_cvttps_pi32(a)
}

func MM_Cvttsssi32(a C.__m128) int {
	return (int)(C._mm_cvttss_si32(a))
}

func MM_Cvttsssi64(a C.__m128) int64 {
	return (int64)(C._mm_cvttss_si64(a))
}

func MM_Divps(a C.__m128, b C.__m128) C.__m128 {
	return C._mm_div_ps(a, b)
}

func MM_Divss(a C.__m128, b C.__m128) C.__m128 {
	return C._mm_div_ss(a, b)
}
/*
func MM_Extractpi16(a C.__m64, b int) int {
	return int(C._mm_extract_pi16(a, C.int(b)))
} You need 2bit integer for b
*/
func MM_GET_EXCEPTION_MASK() uint {
	return (uint)(C._MM_GET_EXCEPTION_MASK())
}

func MM_GET_EXCEPTION_STATE() uint {
	return (uint)(C._MM_GET_EXCEPTION_STATE())
}

func MM_GET_FLUSH_ZERO_MODE() uint {
	return (uint)(C._MM_GET_FLUSH_ZERO_MODE())
}

func MM_GET_ROUNDING_MODE() uint {
	return (uint)(C._MM_GET_ROUNDING_MODE())
}

func MM_Getcsr() uint {
	return (uint)(C._mm_getcsr())
}
/*
func MM_Insertpi16(a C.__m64, b int, c int) C.__m64 {
	return C._mm_insert_pi16(a, (C.int)(b), (C.int)(c))
} You need 2-bit integer c
*/
func M_Maskmovq(a C.__m64, b C.__m64, c []byte) {
	C._m_maskmovq(a, b, (*C.char)(unsafe.Pointer(&c[0])))
}

func M_Pavgb(a C.__m64, b C.__m64) C.__m64 {
	return C._m_pavgb(a, b)
}

func M_Pavgw(a C.__m64, b C.__m64) C.__m64 {
	return C._m_pavgw(a, b)
}
/*
func M_Pextrw(a C.__m64, b int) int {
	return (int)(C._m_pextrw(a, C.int(b)))
} You need 2-bit integer b

func M_Pinsrw(a C.__m64, b int, c int) C.__m64 {
	return C._m_pinsrw(a, C.int(b), C.int(c))
} You need 2-bit integer c
*/
func M_Pmaxsw(a C.__m64, b C.__m64) C.__m64 {
	return C._m_pmaxsw(a, b)
}

func M_Pmaxub(a C.__m64, b C.__m64) C.__m64 {
	return C._m_pmaxub(a, b)
}

func M_Pminsw(a C.__m64, b C.__m64) C.__m64 {
	return C._m_pminsw(a, b)
}

func M_Pminub(a C.__m64, b C.__m64) C.__m64 {
	return C._m_pminub(a, b)
}

func M_Pmovmskb(a C.__m64) int {
	return int(C._m_pmovmskb(a))
}

func M_Pmulhuw(a C.__m64, b C.__m64) C.__m64 {
	return C._m_pmulhuw(a, b)
}

func M_Psadbw(a C.__m64, b C.__m64) C.__m64 {
	return C._m_psadbw(a, b)
}

/*
func M_Pshufw(a C.__m64, b C.int) C.__m64 {
	return C._m_pshufw(a, b)
}
*/

func MM_Maskmovesi64(a C.__m64, b C.__m64, c []byte) {
	C._mm_maskmove_si64(a, b, (*C.char)(unsafe.Pointer(&c[0])))
}

func MM_Maxpi16(a C.__m64, b C.__m64) C.__m64 { 
	return C._mm_max_pi16(a, b)
}

func MM_Maxps(a C.__m128, b C.__m128) C.__m128 {
	return C._mm_max_ps(a, b)
}

func MM_Maxpu8(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_max_pu8(a, b)
}

func MM_Maxss(a C.__m128, b C.__m128) C.__m128 {
	return C._mm_max_ss(a, b)
}

func MM_Minpi16(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_min_pi16(a, b)
}

func MM_Minps(a C.__m128, b C.__m128) C.__m128 {
	return C._mm_min_ps(a, b)
}

func MM_Minpu8(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_min_pu8(a, b)
}

func MM_Minss(a C.__m128, b C.__m128) C.__m128 {
	return C._mm_min_ss(a, b)
}

func MM_Movess(a C.__m128, b C.__m128) C.__m128 {
	return C._mm_move_ss(a, b)
}

func MM_Movehlps(a C.__m128, b C.__m128) C.__m128 {
	return C._mm_movehl_ps(a, b)
}
func MM_Movelhps(a C.__m128, b C.__m128) C.__m128 {
	return C._mm_movelh_ps(a, b)
}

func MM_Movemaskpi8(a C.__m64) C.int {
	return C._mm_movemask_pi8(a)
}

func MM_Movemaskps(a C.__m128) C.int {
	return C._mm_movemask_ps(a)
}

func MM_Mulps(a C.__m128, b C.__m128) C.__m128 {
	return C._mm_mul_ps(a, b)
}

func MM_Mulss(a C.__m128, b C.__m128) C.__m128 {
	return C._mm_mul_ss(a, b)
}

func MM_Mulhipu16(a C.__m64, b C.__m64) C.__m64 {
	return C._mm_mulhi_pu16(a, b)
}

func MM_Orps(a C.__m128, b C.__m128) C.__m128 {
	return C._mm_or_ps(a, b)
}
/*
func MM_Prefetch(a []byte,b C.int) {
	C._mm_prefetch((*C.char)(&a[0]), b)
} You nedd 2-bit integer b
*/
func MM_Rcpps(a C.__m128) C.__m128 {
	return C._mm_rcp_ps(a)
}

func MM_Rcpss(a C.__m128) C.__m128 {
	return C._mm_rcp_ss(a)
}

func MM_Rsqrtps(a C.__m128) C.__m128 {
	return C._mm_rsqrt_ps(a)
}

func MM_Rsqrtss(a C.__m128) C.__m128 {
	return C._mm_rsqrt_ss(a)
}

func MM_Sadpu8(a C.__m64, b C.__m64) C.__m64 { 
	return C._mm_sad_pu8(a, b)
}

func MM_Set1ps(a float32) C.__m128 {
	return C._mm_set1_ps((C.float)(a))
}

func MM_SET_EXCEPTION_MASK(a uint32) {
	C._MM_SET_EXCEPTION_MASK((C.uint)(a))
}

func MM_SET_EXCEPTION_STATE(a uint32) {
	C._MM_SET_EXCEPTION_STATE((C.uint)(a))
}

func MM_SET_FLUSH_ZERO_MODE(a uint32) {
	C._MM_SET_FLUSH_ZERO_MODE((C.uint)(a))
}

func MM_Setps(d float32, c float32, b float32, a float32) C.__m128 {
	return C._mm_set_ps((C.float)(d) ,(C.float)(c), (C.float)(b), (C.float)(d))
}

func MM_Setps1(a float32) C.__m128 {
	return C._mm_set_ps1((C.float)(a))
}

func MM_SET_ROUNDING_MODE(a uint){
	C._MM_SET_ROUNDING_MODE((C.uint)(a))
}

func MM_Setss(a float32) C.__m128 {
	return C._mm_set_ss((C.float)(a))
}

func MM_Setcsr(a uint) {
	C._mm_setcsr((C.uint)(a))
}

func MM_Setrps(d float32, c float32, b float32, a float32) C.__m128 {
	return C._mm_setr_ps((C.float)(d), (C.float)(c), (C.float)(b), (C.float)(a))
}
func MM_Sfence(){
	C._mm_sfence()
}
/*
func MM_Shufflepi16(a C.__m64, b int) C.__m64 {
	return C._mm_shuffle_pi16(a, (C.int)(b))
} // You need a 8-bit integer for b
*/
func MM_Sqrtps(a C.__m128) C.__m128 {
	return C._mm_sqrt_ps(a)
}

func MM_Sqrtss(a C.__m128) C.__m128 {
	return C._mm_sqrt_ss(a)
}

func MM_Streampi(a *C.__m64, b C.__m64) {
	C._mm_stream_pi(a, b)
}

func MM_Streamps(a *float32, b C.__m128){
	C._mm_stream_ps((*C.float)(a), b)
}

func MM_Subps(a C.__m128, b C.__m128) C.__m128 {
	return C._mm_sub_ps(a, b)
}

func MM_Subss(a C.__m128, b C.__m128) C.__m128 {
	return C._mm_sub_ss(a, b)
}

func MM_TRANSPOSE4_PS(a C.__m128, b C.__m128, c C.__m128, d C.__m128){
	C.__MM_TRANSPOSE4_PS(a, b, c, d)
}

func MM_Ucomieqss(a C.__m128, b C.__m128) C.int {
	return C._mm_ucomieq_ss(a, b)
}

func MM_Ucomigess(a C.__m128, b C.__m128) C.int {
	return C._mm_ucomige_ss(a, b)
}

func MM_Ucomigtss(a C.__m128, b C.__m128) C.int {
	return C._mm_ucomigt_ss(a, b)
}

func MM_Ucomiless(a C.__m128, b C.__m128) C.int {
	return C._mm_ucomile_ss(a, b)
}

func MM_Ucomiltss(a C.__m128, b C.__m128) C.int {
	return C._mm_ucomilt_ss(a, b)
}

func MM_Ucomineqss(a C.__m128, b C.__m128) C.int {
	return C._mm_ucomineq_ss(a, b)
}

func MM_Unpackhips(a C.__m128, b C.__m128) C.__m128 {
	return C._mm_unpackhi_ps(a, b)
}

func MM_Unpacklops(a C.__m128, b C.__m128) C.__m128 {
	return C._mm_unpacklo_ps(a, b)
}

func MM_Xorps(a C.__m128, b C.__m128) C.__m128 {
	return C._mm_xor_ps(a, b)
}

