-- лаб 7

-- 1
select emp.last_name, emp.first_name from
employees emp, departments dept
where emp.department_id = dept.department_id
and dept.department_name = 'Sales';

-- 2
select job.job_title, job.min_salary, job.max_salary from
jobs job;

-- 3
select emp.first_name, emp.last_name, emp.salary from
employees emp
where emp.salary > 10000;

-- 4
select dept.department_name, count(emp.employee_id)
from employees emp, departments dept
where emp.department_id = dept.department_id
group by dept.department_name;

-- 5
select emp.first_name, emp.last_name from
employees emp
where to_char(emp.hire_date, 'YYYY') = '2005';

-- 6
select job.job_title, avg(emp.salary)
from employees emp, jobs job
where emp.job_id = job.job_id
group by job.job_title;

-- 7
select emp.first_name, emp.last_name, mgr.first_name as mgr_first_name, mgr.last_name as mgr_last_name
from employees emp, employees mgr
where emp.manager_id = mgr.employee_id;

-- 8
select emp.first_name, emp.last_name, emp.salary, dept.department_name
from employees emp, departments dept
where emp.department_id = dept.department_id
and emp.salary = (select max(salary) from employees where department_id = dept.department_id);

-- 9
select dept.department_name
from employees emp, departments dept, locations loc, countries ctry
where emp.department_id = dept.department_id
and dept.location_id = loc.location_id
and loc.country_id = ctry.country_id
group by dept.department_name
having count(distinct ctry.country_id) > 1;

-- 10
select emp.first_name, emp.last_name, emp.salary, mgr.salary as mgr_salary
from employees emp, employees mgr
where emp.manager_id = mgr.employee_id
and emp.salary > mgr.salary;

-- 11
select dept.department_name, sum(emp.salary) as total_salary
from employees emp, departments dept
where emp.department_id = dept.department_id
group by dept.department_name
order by total_salary desc
where rownum = 1;

-- 12
select emp.first_name, emp.last_name
from employees emp, job_history jh
where emp.employee_id = jh.employee_id
group by emp.employee_id, emp.first_name, emp.last_name
having count(distinct jh.department_id) > 1;

-- 13
select emp.first_name, emp.last_name
from employees emp, employees mgr
where emp.manager_id = mgr.employee_id
and to_char(emp.hire_date, 'MM-YYYY') = to_char(mgr.hire_date, 'MM-YYYY');

-- 14
select emp.first_name, emp.last_name
from employees emp, jobs job
where emp.job_id = job.job_id
and emp.salary = job.max_salary;

-- 15
select dept.department_name
from employees emp, departments dept
where emp.department_id = dept.department_id
group by dept.department_name
having avg(emp.salary) > (select avg(salary) from employees);

-- 16
select dept.department_name, max(emp.salary) - min(emp.salary) as salary_diff
from employees emp, departments dept
where emp.department_id = dept.department_id
group by dept.department_name
order by salary_diff desc
where rownum = 1;

-- 17
select distinct jh.employee_id, emp.first_name, emp.last_name
from job_history jh, employees emp
where jh.employee_id = emp.employee_id
and exists (select 1 from job_history jh2
where jh2.employee_id = jh.employee_id
and jh2.job_id = jh.job_id
and jh2.department_id <> jh.department_id);

-- 18
select emp.first_name, emp.last_name
from employees emp
where not exists (select 1 from jobs job
where job.job_id in (select distinct job_id from employees where department_id = emp.department_id)
and job.job_id not in (select job_id from job_history where employee_id = emp.employee_id)
and job.job_id <> emp.job_id);

-- 19
select emp.first_name, emp.last_name
from employees emp, jobs job
where emp.job_id = job.job_id
and emp.salary > (select avg(salary) from employees where job_id = emp.job_id);

-- 20
select dept.department_name
from employees emp, departments dept
where emp.department_id = dept.department_id
group by dept.department_name
having count(distinct to_char(hire_date, 'YYYY')) = 1;

-- 21
select emp.first_name, emp.last_name
from employees emp, employees mgr
where emp.manager_id = mgr.employee_id
and emp.hire_date < mgr.hire_date;

-- 22
select emp.first_name, emp.last_name
from employees emp, job_history jh
where emp.employee_id = jh.employee_id
group by emp.employee_id, emp.first_name, emp.last_name
having count(distinct jh.job_id) > 2;

-- 23
select dept.department_name
from employees emp, departments dept, job_history jh
where emp.department_id = dept.department_id
and emp.employee_id = jh.employee_id
group by dept.department_name
having count(distinct jh.job_id) > 1;


-- 24
select emp.first_name, emp.last_name
from employees emp
where not exists (select 1 from job_history jh 
where jh.employee_id = emp.employee_id and jh.department_id <> emp.department_id);

-- 25
select emp.first_name, emp.last_name
from employees emp, jobs job, departments dept
where emp.job_id = job.job_id
and emp.department_id = dept.department_id
and emp.salary = job.max_salary
and emp.salary =  (select max(salary) from employees where department_id = dept.department_id);


-- лаб 8

-- 1
select emp.last_name, dept.department_name from
employee emp, department dept
where emp.deptartment_id = dept.deptartment_id;

-- 2
select * from
employee emp, department dept
where emp.deptartment_id = 60 and dept.deptartment_id = 60;

-- 3
select emp.employee_id, emp.manager from
employee emp;

-- 4
select dept.department_name, count(emp.employee_id) 
from employee emp, department dept
where emp.deptartment_id = dept.deptartment_id
group by dept.department_name;

-- 5
select dept.department_name, count(emp.employee_id) 
from employee emp, department dept
where emp.deptartment_id = dept.deptartment_id
group by dept.department_name
having count(emp.employee_id) > 5;

-- 6
select dept.department_name, avg(emp.salary) 
from employee emp, department dept
where emp.deptartment_id = dept.deptartment_id
group by dept.department_name;


-- 7
select emp.employee_id
from employee emp, (select select dept.department_id d_id, avg(emp.salary) d_sal
from employee emp, department dept
where emp.deptartment_id = dept.deptartment_id
group by dept.department_id;) dept_avg_sal
where emp.department_id = dept_avg_sal.d_id and emp.salary > dept_avg_sal.d_sal;

-- 8
select *
from employee emp
where emp.manager = null;

-- 9
select mgr.last_name as manager, count(emp.employee_id) as subordinates from
employees emp, employees mgr
where emp.manager_id = mgr.employee_id
group by mgr.last_name;

-- 10
select unique emp.job
from employee emp;

-- 11
select dept.department_name, max(emp.salary) as max_salary from
employees emp, departments dept
where emp.department_id = dept.department_id
group by dept.department_name;

-- 12
select emp.last_name, job.job_title, loc.city from
employees emp, jobs job, departments dept, locations loc
where emp.job_id = job.job_id
and emp.department_id = dept.department_id
and dept.location_id = loc.location_id;

-- 13
select emp.last_name, emp.salary, job.job_title from
employees emp, jobs job
where emp.job_id = job.job_id
and emp.salary < (select avg(salary) from employees where job_id = emp.job_id);

-- 14
select emp.last_name as employee, emp.hire_date as emp_hire_date,
mgr.last_name as manager, mgr.hire_date as mgr_hire_date from
employees emp, employees mgr
where emp.manager_id = mgr.employee_id;

-- 15
select emp.last_name as employee, sub.last_name as subordinate from
employees emp, employees sub
where emp.employee_id = sub.manager_id(+)
order by emp.last_name;


-- усложненные

-- 1
select dept.department_name, avg(emp.salary) as avg_salary from
employees emp, departments dept
where emp.department_id = dept.department_id
group by dept.department_name
having avg(emp.salary) > 500;

-- 2
select emp.last_name, emp.salary from
employees emp, departments dept, jobs job
where emp.department_id = dept.department_id
and emp.job_id = job.job_id
and emp.salary > (select avg(salary) from employees where department_id = dept.department_id)
and emp.salary > (select avg(salary) from employees where job_id = job.job_id);

-- 3
select emp.last_name, job.job_title, emp.salary from
employees emp, employees mgr, jobs job
where emp.manager_id = mgr.employee_id
and emp.job_id = job.job_id
and mgr.last_name = 'King';

-- 4
select emp.last_name, job.job_title, dept.department_name, emp.hire_date from
employees emp, jobs job, departments dept
where emp.job_id = job.job_id
and emp.department_id = dept.department_id
and emp.hire_date > add_months(sysdate, -6);

-- 5
select dept.department_name, max(emp.salary) as max_salary, min(emp.salary) as min_salary from
employees emp, departments dept
where emp.department_id = dept.department_id
group by dept.department_name;

-- 6
select dept.department_name, count(emp.employee_id) as employee_count,
(select count(*) from employees) as total_employees from
employees emp, departments dept
where emp.department_id = dept.department_id
group by dept.department_name;

-- 7
select emp1.last_name, emp1.salary, job.job_title from
employees emp1, jobs job
where emp1.job_id = job.job_id
and exists (select 1 from employees emp2
where emp2.job_id = emp1.job_id
and emp2.employee_id != emp1.employee_id
and emp2.salary = emp1.salary);

-- 8
select emp.last_name as employee, emp.salary as emp_salary,
mgr.last_name as manager, mgr.salary as mgr_salary from
employees emp, employees mgr
where emp.manager_id = mgr.employee_id
and abs(emp.salary - mgr.salary) > 1000;

-- 9
select mgr.last_name as manager, dept.department_name from
employees mgr, departments dept
where mgr.department_id = dept.department_id
and (select count(*) from employees emp where emp.manager_id = mgr.employee_id) > 1;

-- 10
select dept.department_name from
employees emp, departments dept
where emp.department_id = dept.department_id
and emp.salary = (select max(salary) from employees);


-- лаб 9

-- 1
select emp.employee_id
from HR.employees emp
join (select dept.department_id, avg(emp1.salary) d_sal
from HR.employees emp1 
join HR.departments dept on emp1.department_id = dept.department_id
group by dept.department_id) dept_avg_sal
on emp.department_id = dept_avg_sal.department_id
where emp.salary > dept_avg_sal.d_sal;

-- 2
select *
from HR.employees
where department_id = (
    select department_id
    from hr.departments
    group by department_id
    order by count(*) desc
    fetch first 1 row only
);

-- 3
select *
from hr.employees
where salary > (
    select salary
    from hr.employees
    where employee_id = 101
    fetch first 1 row only
);

-- 4
select * 
from hr.employees
where department_id = (
    select department_id
    from hr.employees
    where employee_id = 102
    fetch first 1 row only
);

-- 5
select *
from hr.employees
where salary > (
    select max(salary)
    from hr.employees
    where department_id = 30
);

-- 6
select *
from hr.employees emp
join (
    select job_id, avg(salary) avg_salary
    from hr.employees
    group by job_id
) avg_sal on emp.job_id = avg_sal.job_id
where emp.salary < avg_sal.avg_salary;

-- 7
select *
from hr.employees
where job_id = 'IT_PROG'
and salary > (
    select max(salary)
    from hr.employees
    where job_id = 'SAMAN'
);

-- 8
select emp.*
from hr.employees emp
where department_id in (
    select department_id
    from hr.employees
    group by department_id
    having min(salary) > 3000
);

-- 9
select *
from hr.employees
where hire_date > (
    select max(hire_date)
    from hr.employees
    where job_id = 'HR_REP'
);

-- 10
select *
from hr.employees
where salary < (
    select avg(salary)
    from hr.employees
);

-- 11
select department_id
from hr.employees
group by department_id
having avg(salary) > 5000;

-- 12
select department_id
from hr.employees
group by department_id
having min(salary) > 2000;

-- 13
select distinct department_id
from hr.employees emp
where salary > (
    select avg(salary)
    from hr.employees
    where job_id = emp.job_id
);

-- 14
select department_id
from hr.employees
group by department_id
having max(salary) > 10000;

-- 15
select department_id
from hr.employees
group by department_id
having count(*) > (
    select avg(count(*))
    from hr.employees
    group by department_id
);

-- 16
select distinct mgr.employee_id
from hr.employees emp
join hr.employees mgr on emp.manager_id = mgr.employee_id;

-- 17
select distinct mgr.employee_id
from hr.employees emp
join hr.employees mgr on emp.manager_id = mgr.employee_id
where mgr.department_id = 50;

-- 18
select mgr.*
from hr.employees mgr
where mgr.employee_id in (
    select manager_id
    from hr.employees
    where manager_id is not null
)
and mgr.salary > (
    select avg(salary)
    from hr.employees
    where department_id = mgr.department_id
);

-- 19
select emp.*
from hr.employees emp
where department_id in (
    select department_id
    from hr.employees
    group by department_id
    having avg(salary) > 6000
);

-- 20
select *
from hr.employees
where job_id = 'SAREP'
and department_id in (
    select distinct department_id
    from hr.employees
    where job_id = 'ITPROG'
);


-- дз

-- 1
select emp.*
from hr.employees emp
join (
    select department_id, avg(salary) avg_sal
    from hr.employees
    group by department_id
) dept_avg on emp.department_id = dept_avg.department_id
where emp.salary > dept_avg.avg_sal
and emp.hire_date > '31.12.2015';

-- 2
select emp.*
from hr.employees emp
where department_id = (
    select department_id
    from hr.employees
    group by department_id
    order by avg(salary) desc
    fetch first 1 row only
);

-- 3
select emp.*
from hr.employees emp
join (
    select job_id, max(salary) max_sal
    from hr.employees
    group by job_id
) job_max on emp.job_id = job_max.job_id
where emp.salary > job_max.max_sal * 0.1;

-- 4
select mgr.*
from hr.employees mgr
where mgr.employee_id in (
    select distinct manager_id
    from hr.employees
    where manager_id is not null
)
and mgr.salary > (
    select avg(salary)
    from hr.employees
    where manager_id = mgr.employee_id
);

-- 5
select emp.*
from hr.employees emp
where department_id in (
    select distinct department_id
    from hr.employees
    where job_id = 'HR_REP'
)
and salary > 5000;

-- 6
select department_id
from hr.employees
group by department_id
having avg(salary) > (
    select avg(salary)
    from hr.employees
);

-- 7
select department_id
from hr.employees
group by department_id
having count(*) > (
    select avg(count(*))
    from hr.employees
    group by department_id
);

-- 8
select department_id
from hr.employees
group by department_id
having max(salary) > (
    select max(salary)
    from hr.employees
    where department_id = 50
);

-- 9
select distinct mgr.*
from hr.employees mgr
join hr.employees emp on mgr.employee_id = emp.manager_id
where emp.salary > (
    select avg(salary)
    from hr.employees
);

-- 10
select distinct mgr.*
from hr.employees mgr
join hr.employees emp on mgr.employee_id = emp.manager_id
where mgr.department_id = 60
and emp.salary > (
    select avg(salary)
    from hr.employees
    where job_id = emp.job_id
);

-- 11
select emp.*
from hr.employees emp
where department_id in (
    select department_id
    from hr.employees
    group by department_id
    having min(salary) > 3000
);

-- 12
select emp.*
from hr.employees emp
where job_id = 'SAREP'
and department_id in (
    select distinct department_id
    from hr.employees
    where job_id = 'ITPROG'
)
and salary < (
    select avg(salary)
    from hr.employees
    where job_id = 'SA_REP'
);

-- 13
select 
    department_id,
    avg(salary) avg_sal
from hr.employees
group by department_id
order by avg_sal desc;

-- 14
select job_id, count(*)
from hr.employees emp 
join (
    select job_id, max(salary) max_sal
    from hr.employees
    group by job_id
) job_max_sal on emp.job_id = job_max_sal.job_id 
and emp.sal = job_max_sal.max_sal
group by job_id;

-- 15
???

-- 16
select emp.*
from hr.employees emp
join (
    select job_id, max(salary) max_sal
    from hr.employees
    group by job_id
) job_max on emp.job_id = job_max.job_id and emp.salary = job_max.max_sal;

-- 17
select emp.*
from hr.employees emp
join (
    select department_id, min(salary) min_sal
    from hr.employees
    group by department_id
) dept_min on emp.department_id = dept_min.department_id and emp.salary = dept_min.min_sal;

-- 18
???

-- 19
select 
    emp.*,
    job_count.emp_count
from hr.employees emp
join (
    select job_id, count(*) emp_count
    from hr.employees
    group by job_id
) job_count on emp.job_id = job_count.job_id;

-- 20
select 
    emp.*,
    dept_count.emp_count,
    dept_max.max_sal
from hr.employees emp
join (
    select department_id, count(*) emp_count, max(salary) max_sal
    from hr.employees
    group by department_id
) dept_count on emp.department_id = dept_count.department_id;

-- 15, 18 - ?


-- лаб 14

-- легкий

-- 1
create or replace function get_job_title(
    p_job_id in varchar2
) return varchar2 as
    v_job_title varchar2(35);
begin
    select job_title into v_job_title
    from jobs
    where job_id = p_job_id;
    
    return v_job_title;
end;


-- 2
create or replace procedure add_department(
    p_department_name in varchar2,
    p_location_id in number,
    p_manager_id in number default null
) as
    v_department_id number;
begin
    select nvl(max(department_id), 0) + 10 into v_department_id
    from departments;
    
    insert into departments (department_id, department_name, location_id, manager_id)
    values (v_department_id, p_department_name, p_location_id, p_manager_id);
end;


-- 3
create or replace function is_manager(
    p_employee_id in number
) return boolean as
    v_count number;
begin
    select count(*) into v_count
    from employees
    where manager_id = p_employee_id;
    
    return v_count > 0;
end;


-- 4
create or replace procedure update_employee_email(
    p_employee_id in number,
    p_new_email in varchar2
) as
begin
    update employees
    set email = p_new_email
    where employee_id = p_employee_id;
end;


-- 5
create or replace function get_employee_count
return number as
    v_count number;
begin
    select count(*) into v_count
    from employees;
    
    return v_count;
end;

-- средний

-- 6
create or replace procedure increase_sal (
    p_proc number;
) as
begin
    update employees
    set salary = salary * (1 + p_proc / 100);
end;

--7
create or replace function max_salary (
    p_dept_id in number
) return number as 
    v_max_sal number;
begin
    select max(salary) into v_max_sal
    from employees 
    where department_id = p_dept_id;

    return v_max_sal;
end;

-- 8
create or replace procedure delete_dept(
    p_dept_id in number
) as
begin
    delete from departments
    where department_id = p_dept_id;
end;

-- 9
declare
begin
    for emp in (select * from employees) loop
        dbms_output.put_line(emp.last_name || ': ' || emp.salary);
    end loop;
end;

-- 10
create or replace function get_dept_employees(
    p_dept_id in number
) return sys_refcursor as 
    v_cursor sys_refcursor;
begin
    open v_cursor for
        select employee_id, last_name, first_name
        from employees
        where department_id = p_dept_id;
    
    return v_cursor;
end;

-- высокий

-- 11
create or replace procedure change_dept(
    p_from_dept_id in number,
    p_to_dept_id in number
) as
    v_count number;
begin
    update employees
    set department_id = p_to_dept_id
    where department_id = p_from_dept_id;
end;


-- 12
create or replace function count_emp_by_salary(
    p_salary in number
) return number as
    v_count number;
begin
    select count(*) into v_count
    from employees
    where salary > p_salary;
    
    return v_count;
end;


-- 13
create or replace procedure set_dept_manager(
    p_dept_id in number,
    p_new_manager_id in number
) as
    v_count number;
begin
    update employees
    set manager_id = p_new_manager_id
    where department_id = p_dept_id
    and employee_id != p_new_manager_id;
end;


-- 14
declare
    cursor c_emp is
        
        for update;
begin
    for emp in (select employee_id, salary from employees) loop
        if emp.salary > 20000 then
            update employees
            set job_id = 'JOB1'
            where emp.employee_id = employee_id;
        elsif emp.salary > 10000 then
            update employees
            set job_id = 'JOB2'
            where emp.employee_id = employee_id;
        end if;
    end loop;
end;


-- 15
create or replace function get_avg_age
return number as
    v_avg_age number;
begin
    select avg(months_between(sysdate, hire_date)/12) into v_avg_age
    from employees;
    
    return v_avg_age;
end;

-- экспертный

-- 16
create or replace procedure update_job as
begin
    for emp_rec in (select employee_id, months_between(sysdate, hire_date) as exp
        from employees) loop
        if emp_rec.exp >= 24 then
            update employees
            set job_id = 'good'
            where emp.employee_id = employee_id;
        elsif emp_rec.exp >= 12 then
            update employees
            set job_id = 'notgood'
            where emp.employee_id = employee_id;
        end if;
    end loop;
end;


-- 17
create or replace function get_podchin(
    p_manager_id in number
) return sys_refcursor as
    v_cursor sys_refcursor;
begin
    open v_cursor for
        select employee_id, last_name, first_name, job_id
        from employees
        where manager_id = p_manager_id;
    return v_cursor;
end;


-- 18
declare
begin
    for dept in (
        select d.department_id, count(e.employee_id) as emp_count
        from departments d
        join employees e on d.department_id = e.department_id
        group by d.department_id
        order by emp_count
    ) loop
        dbms_output.put_line(dept.department_id || ': ' || dept.emp_count);
    end loop;
end;


-- 19
create or replace procedure delete_emp_without_manager as
begin
    delete from employees
    where manager_id is null;
end;


-- 20
create or replace function get_unique_jobs
return sys_refcursor as
    v_cursor sys_refcursor;
begin
    open v_cursor for
        select distinct job_id, job_title
        from jobs
        order by job_title;
    return v_cursor;
end;