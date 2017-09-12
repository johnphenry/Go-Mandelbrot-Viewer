package mandel
import("math/cmplx")
func Calc_pixel( x, y, w, h int, offset complex128, scale complex128 )(uint8,uint8,uint8){
	c:= func (x,y,w,h int)complex128{
		axis:=func(a,d int)float64{
			return (float64( a ) / float64( d ) ) * 4 -2
		}
		aspect := float64( h ) / float64( w )
		return complex( axis( x, w ), axis( y, h )* aspect )*scale + offset
	}( x,y,w,h )
	z:=c
	i := 0
	const LIMIT = 32
	for ; i < LIMIT && cmplx.Abs(z) < 4; i++ {
		z = z*z + c
	}
	if i == LIMIT{
		return 0,0,0
	}
	col:= uint8((float64(i) / float64(LIMIT)) * 255.0)
	return col,col,col
}
