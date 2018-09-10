@extends('layouts.project')

@section('project')

            <div class="card">
                <div class="card-header">Evaluation</div>
                <div class="card-body">
                	<form>
						<div class="input-group">
						  <div class="input-group-prepend">
						    <span class="input-group-text">기여 내용</span>
						  </div>
						  <textarea class="form-control" aria-label="evaluation"></textarea>
						</div>
						<btn class="btn btn-secondary mt-3">제출</btn>
					</form>
                </div>
            </div>
@endsection