package corelib

type VoxelData struct{
	Width, Height, Depth int
	Data []byte
}

func NewVoxelData(width, height, depth int) *VoxelData{
	return &VoxelData{Width:width, Height:height, Depth:depth,Data:make([]byte, width*height*depth)}
}

func (v*VoxelData)Getf(x,y,z int) float32{
	if x>v.Width || x<0 {
		return 0
	}
	if y>v.Height || y<0 {
		return 0
	}
	if z>v.Depth || z < 0{
		return 0
	}
	return float32(v.Data[x+y*v.Width+z*v.Height*v.Width] / 255.0)
}


func (v*VoxelData)Get3f(x,y,z int) float32{
	if x>v.Width || x<0 {
		return 0
	}
	if y>v.Height || y<0 {
		return 0
	}
	if z>v.Depth || z < 0{
		return 0
	}
	var agg int = int(v.Data[x+y*v.Width+z*v.Height*v.Width])
	var n int = 1

	if x < v.Width-1 {
		n+=1
		agg += int(v.Data[x+1+y*v.Width+z*v.Height*v.Width])

		if y < v.Height-1 {
			n+=1
			agg += int(v.Data[x+1+(y+1)*v.Width+z*v.Height*v.Width])
		}

		if y > 0{
			n+=1
			agg += int(v.Data[x+1+(y-1)*v.Width+z*v.Height*v.Width])
		}

		if z < v.Depth-1 {
			n+=1
			agg += int(v.Data[x+1+y*v.Width+(z+1)*v.Height*v.Width])
		}

		if z > 0{
			n+=1
			agg += int(v.Data[x+1+y*v.Width+(z-1)*v.Height*v.Width])
		}
	}

	if x > 0{
		n+=1
		agg += int(v.Data[x-1+y*v.Width+z*v.Height*v.Width])
		if y < v.Height-1 {
			n+=1
			agg += int(v.Data[x-1+(y+1)*v.Width+z*v.Height*v.Width])
		}

		if y > 0{
			n+=1
			agg += int(v.Data[x-1+(y-1)*v.Width+z*v.Height*v.Width])
		}

		if z < v.Depth-1 {
			n+=1
			agg += int(v.Data[x-1+y*v.Width+(z+1)*v.Height*v.Width])
		}

		if z > 0{
			n+=1
			agg += int(v.Data[x-1+y*v.Width+(z-1)*v.Height*v.Width])
		}
	}

	if y < v.Height-1 {
		n+=1
		agg += int(v.Data[x+(y+1)*v.Width+z*v.Height*v.Width])

		if x < v.Width-1 {
			n+=1
			agg += int(v.Data[x+1+(y+1)*v.Width+z*v.Height*v.Width])
		}

		if x > 0{
			n+=1
			agg += int(v.Data[x-1+(1+y)*v.Width+z*v.Height*v.Width])
		}

		if z < v.Depth-1 {
			n+=1
			agg += int(v.Data[x+(1+y)*v.Width+(z+1)*v.Height*v.Width])
		}

		if z > 0{
			n+=1
			agg += int(v.Data[x+(1+y)*v.Width+(z-1)*v.Height*v.Width])
		}
	}

	if y > 0{
		n+=1
		agg += int(v.Data[x+(y-1)*v.Width+z*v.Height*v.Width])

		if x < v.Width-1 {
			n+=1
			agg += int(v.Data[x+1+(y-1)*v.Width+z*v.Height*v.Width])
		}

		if x > 0{
			n+=1
			agg += int(v.Data[x-1+(y-1)*v.Width+z*v.Height*v.Width])
		}

		if z < v.Depth-1 {
			n+=1
			agg += int(v.Data[x+(y-1)*v.Width+(z+1)*v.Height*v.Width])
		}

		if z > 0{
			n+=1
			agg += int(v.Data[x+(y-1)*v.Width+(z-1)*v.Height*v.Width])
		}
	}

	if z < v.Depth-1 {
		n+=1
		agg += int(v.Data[x+y*v.Width+(z+1)*v.Height*v.Width])
		if x < v.Width-1 {
			n+=1
			agg += int(v.Data[x+1+y*v.Width+(1+z)*v.Height*v.Width])
		}

		if x > 0{
			n+=1
			agg += int(v.Data[x-1+y*v.Width+(1+z)*v.Height*v.Width])
		}

		if y < v.Height-1 {
			n+=1
			agg += int(v.Data[x+(y+1)*v.Width+(1+z)*v.Height*v.Width])
		}

		if y > 0{
			n+=1
			agg += int(v.Data[x+(y-1)*v.Width+(1+z)*v.Height*v.Width])
		}
	}

	if z > 0{
		n+=1
		agg += int(v.Data[x+y*v.Width+(z-1)*v.Height*v.Width])
		if x < v.Width-1 {
			n+=1
			agg += int(v.Data[x+1+y*v.Width+(z-1)*v.Height*v.Width])
		}

		if x > 0{
			n+=1
			agg += int(v.Data[x-1+y*v.Width+(z-1)*v.Height*v.Width])
		}

		if y < v.Height-1 {
			n+=1
			agg += int(v.Data[x+(y+1)*v.Width+(z-1)*v.Height*v.Width])
		}

		if y > 0{
			n+=1
			agg += int(v.Data[x+(y-1)*v.Width+(z-1)*v.Height*v.Width])
		}
	}


	return float32(agg / (255.0*n))
}

func (v*VoxelData)Setf(x,y,z int, value float32){
	if x>v.Width || x<0 {
		return
	}
	if y>v.Height || y<0 {
		return
	}
	if z>v.Depth || z < 0{
		return
	}
	v.Data[x+y*v.Width+z*v.Height*v.Width]  = byte(value*255.0);
}

func (v*VoxelData)Set(x,y,z int, value byte){
	if x>v.Width || x<0 {
		return
	}
	if y>v.Height || y<0 {
		return
	}
	if z>v.Depth || z < 0{
		return
	}
	v.Data[x+y*v.Width+z*v.Height*v.Width]  = value;
}